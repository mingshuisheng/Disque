package logic

import (
	"disqueBackend/dao"
	"disqueBackend/models"
	"disqueBackend/utils/fileUtils"
	"disqueBackend/utils/zipUtils"
	"emperror.dev/errors"
	"github.com/google/uuid"
	cp "github.com/otiai10/copy"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

func GetFileList(parentID uint) []models.File {
	fileList, err := dao.QueryFileList(parentID)
	if err != nil {
		return []models.File{}
	}
	return fileList
}

func MakeDir(parentID uint, name string) (models.File, error) {
	treeID := "0"
	if parentID != 0 {
		info, err := GetFileInfo(parentID)
		if err != nil {
			return models.File{}, errors.Combine(err, &ParentFileNotExitError{}) //errors.Wrapf(err, "找不到文件夹")
		}
		treeID = info.TreeID
	}

	oldFile, err := dao.QueryFileByParentIDAndName(parentID, name)

	if err != nil {
		return models.File{}, errors.Wrapf(err, "数据库错误")
	}

	if oldFile.ID != 0 {
		return oldFile, nil
	}

	newDir := models.File{ParentID: parentID, Name: name, IsDir: true}
	err = dao.InsertFile(&newDir)
	if err != nil {
		return models.File{}, errors.Combine(err, &FolderCreateError{})
	}
	newDir.TreeID = treeID + "-" + strconv.Itoa(int(newDir.ID))
	err = dao.UpdateFile(&newDir)

	if err != nil {
		return models.File{}, errors.Combine(err, &FolderCreateError{})
	}
	return newDir, nil
}

func MakeAllDir(parentID uint, path string) (models.File, error) {
	folderNames := strings.Split(path, "/")
	var lastFile models.File
	var err error

	for _, folderName := range folderNames {
		if folderName == "" {
			continue
		}

		lastFile, err = MakeDir(parentID, folderName)
		if err != nil {
			return models.File{}, errors.Wrapf(err, "创建文件夹%s失败", folderName)
		}
		parentID = lastFile.ID
	}
	return lastFile, nil
}

func GetAllParentFileList(ID uint) []models.File {
	fileList, err := dao.QueryAllParentList(ID)
	if err != nil {
		return []models.File{}
	}

	return fileList
}

func GetFileInfo(ID uint) (models.File, error) {
	return dao.QueryFile(ID)
}

func SaveFile(parentID uint, file *multipart.FileHeader, fullPath string) (err error) {

	//将文件保存到本地
	fileName, dst, ext, err := fileUtils.SaveFile(file)
	if err != nil {
		return errors.Wrapf(err, "文件保存失败")
	}

	//查看是否需要创建文件夹，并且重置父文件夹id
	if fullPath != "" {
		start := 0
		if strings.Index(fullPath, "/") == 0 {
			start = 1
		}
		fullPath = fullPath[start:strings.LastIndex(fullPath, "/"+file.Filename)]
		if fullPath != "" {
			dir, err := MakeAllDir(parentID, fullPath)
			if err != nil {
				return errors.Combine(err)
			}
			parentID = dir.ID
		}
	}

	//本地文件映射
	localFileInfo := models.LocalFile{
		LocalPath: dst,
	}

	err = dao.InsertLocalFile(&localFileInfo)

	if err != nil {
		return err
	}

	oldFile, err := dao.QueryFileByParentIDAndName(parentID, file.Filename)
	if err != nil {
		return errors.Wrapf(err, "文件保存失败")
	}

	if oldFile.ID != 0 {
		//更新
		oldFile.LocalFileID = localFileInfo.ID
		err = dao.UpdateFile(&oldFile)
	} else {
		//创建
		fileInfo := models.File{
			Name:        fileName,
			ParentID:    parentID,
			IsDir:       false,
			LocalFileID: localFileInfo.ID,
			ExtType:     ext,
		}
		err = dao.InsertFile(&fileInfo)

		if err != nil {
			return errors.Wrapf(err, "文件保存失败")
		}

		//计算新的TreeId
		treeID := "0"
		//校验parentID是否有效
		if parentID != 0 {
			info, err := GetFileInfo(parentID)
			if err != nil || info.ID == 0 || info.IsDir == false {
				return errors.Combine(err, &ParentFileNotExitError{})
			}
			treeID = info.TreeID
		}

		fileInfo.TreeID = treeID + "-" + strconv.Itoa(int(fileInfo.ID))

		err = dao.UpdateFile(&fileInfo)
	}

	return errors.Wrapf(err, "文件保存失败")
}

func RenameFile(ID uint, newFileName string) error {
	fileInfo, err := GetFileInfo(ID)

	if err != nil {
		return errors.Combine(err, &FileNotExitError{})
	}

	fileInfo.Name = newFileName
	err = dao.UpdateFile(&fileInfo)
	return errors.Wrapf(err, "重命名失败")
}

func DeleteFile(ID uint) error {
	return errors.Wrapf(dao.DeleteFile(ID), "文件删除失败")
}

func GetFileLocalPathAndFileName(ID uint) (path string, fileName string, err error) {
	path = ""
	fileName = ""
	file, err := dao.QueryFile(ID)
	if err != nil {
		err = errors.Combine(err, &FileNotExitError{})
		return
	}

	if file.IsDir {
		fileName = file.Name + ".zip"
		path, err = toZipFile(file)
		err = errors.Wrapf(err, "zip文件创建失败")
		return
	}

	fileName = file.Name
	path, err = getSingleFile(file.LocalFileID)
	err = errors.Wrapf(err, "文件实体丢失")
	return
}

func toZipFile(file models.File) (path string, err error) {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")

	tempPath := "./temp/" + id

	err = os.MkdirAll(tempPath, os.ModePerm)
	if err != nil {
		return
	}

	files, err := dao.QueryAllChildrenByTreeID(file.TreeID)

	if err != nil {
		return
	}

	files = append(files, file)

	root := buildTree(file.ID, files)

	err = mkDirAndCopyFile(tempPath, root)
	if err != nil {
		return "", err
	}
	err = zipUtils.ZipSource(tempPath+"/"+file.Name, tempPath+"/"+file.Name+".zip")
	if err != nil {
		return "", err
	}

	path = tempPath + "/" + file.Name + ".zip"

	return
}

func mkDirAndCopyFile(bastPath string, root *TreeNode) error {
	if root.FileInfo.IsDir {
		//创建当前目录
		currentPath := bastPath + "/" + root.FileInfo.Name
		err := os.MkdirAll(currentPath, os.ModePerm)
		if err != nil {
			return err
		}
		//递归创建子目录
		if len(root.Children) > 0 {
			for _, child := range root.Children {
				err := mkDirAndCopyFile(currentPath, child)
				if err != nil {
					return err
				}
			}
		}
	} else {
		//如果是文件则复制
		localFileInfo, err := dao.QueryLocalFile(root.FileInfo.LocalFileID)
		if err != nil {
			return err
		}
		err = cp.Copy(localFileInfo.LocalPath, bastPath+"/"+root.FileInfo.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

type TreeNode struct {
	FileInfo models.File
	Children []*TreeNode
}

func buildTree(rootID uint, files []models.File) *TreeNode {
	var root *TreeNode = nil

	var m = make(map[uint]*TreeNode)

	for _, file := range files {
		m[file.ID] = &TreeNode{FileInfo: file}
		if file.ID == rootID {
			root = m[file.ID]
		}
	}

	for _, file := range files {
		if file.ID != rootID {
			current := m[file.ParentID]
			if current.Children == nil {
				current.Children = make([]*TreeNode, 0)
			}
			current.Children = append(current.Children, m[file.ID])
		}
	}

	return root
}

func getSingleFile(ID uint) (path string, err error) {
	localFile, err := dao.QueryLocalFile(ID)
	if err != nil {
		return
	}
	path = localFile.LocalPath
	return
}
