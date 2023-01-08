package logic

import (
	"disqueBackend/dao"
	"disqueBackend/models"
	"disqueBackend/utils/fileUtils"
	"disqueBackend/utils/zipUtils"
	"errors"
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

func MakeDir(parentID uint, name string) error {
	treeID := "0"
	if parentID != 0 {
		info, err := GetFileInfo(parentID)
		if err != nil {
			return err
		}
		treeID = info.TreeID
	}

	newDir := models.File{ParentID: parentID, Name: name, IsDir: true}
	err := dao.InsertFile(&newDir)
	if err != nil {
		return err
	}
	newDir.TreeID = treeID + "-" + strconv.Itoa(int(newDir.ID))
	err = dao.UpdateFile(&newDir)

	if err != nil {
		return err
	}
	return nil
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

func SaveFile(parentID uint, file *multipart.FileHeader) (err error) {

	fileName, dst, ext, err := fileUtils.SaveFile(file)
	if err != nil {
		//log.Println("文件保存失败")
		//ctx.JSON(http.StatusBadRequest, gin.H{"msg": "文件保存失败"})
		return
	}

	treeID := "0"
	//校验parentID是否有效
	if parentID != 0 {
		info, err := GetFileInfo(parentID)
		if err != nil || info.ID == 0 || info.IsDir == false {
			err = errors.New("正在操作的文件不存在")
			return err
		}
		treeID = info.TreeID
	}

	localFileInfo := models.LocalFile{
		LocalPath: dst,
	}

	err = dao.InsertLocalFile(&localFileInfo)

	if err != nil {
		return err
	}

	fileInfo := models.File{
		Name:        fileName,
		ParentID:    parentID,
		IsDir:       false,
		LocalFileID: localFileInfo.ID,
		ExtType:     ext,
	}

	err = dao.InsertFile(&fileInfo)

	if err != nil {
		return err
	}

	fileInfo.TreeID = treeID + "-" + strconv.Itoa(int(fileInfo.ID))

	err = dao.UpdateFile(&fileInfo)

	return err
}

func GetFileLocalPathAndFileName(ID uint) (path string, fileName string, err error) {
	path = ""
	fileName = ""
	file, err := dao.QueryFile(ID)
	if err != nil {
		return
	}

	if file.IsDir {
		fileName = file.Name + ".zip"
		path, err = toZipFile(file)
		return
	}

	fileName = file.Name
	path, err = getSingleFile(file.LocalFileID)
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
