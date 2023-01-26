package service

import (
	"context"
	"disqueBackend/dao"
	"disqueBackend/models"
	"disqueBackend/utils/errorUtils"
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

type FileService struct {
	*_BaseService
	fileDao      *dao.FileDao
	localFileDao *dao.LocalFileDao
}

func CreateFileService(ctx context.Context) *FileService {
	baseService := createBaseService(ctx)
	return &FileService{
		_BaseService: baseService,
		fileDao:      dao.CreateFileDao(baseService.transactionHolder),
		localFileDao: dao.CreateLocalFileDao(baseService.transactionHolder),
	}
}

func (fileService *FileService) GetFileList(ID models.PrimaryKey) ([]models.File, error) {
	fileList, err := fileService.fileDao.ListChildren(ID)
	return fileList, errors.Wrapf(err, "文件不存在")
}

func (fileService *FileService) MakeDir(parentID models.PrimaryKey, name string) (newDir models.File, err error) {
	err = fileService.transaction(func() error {
		fileDao := fileService.fileDao
		treeID := models.RootFile.TreeID
		if parentID != 0 {
			info, err := fileDao.Find(parentID)
			if err != nil {
				return errorUtils.Combine(err, ParentFileNotExitError)
			}
			treeID = info.TreeID
		}

		newDir = models.File{ParentID: parentID, Name: name, IsDir: true}

		oldFile, err := fileDao.FindByParentIDAndName(parentID, name)

		if err != nil {
			return errorUtils.Combine(err, ParentFileNotExitError)
		}

		if oldFile.ID != 0 {
			return nil
		}

		err = fileDao.Insert(&newDir)
		if err != nil {
			return errorUtils.Combine(err, FolderCreateError)
		}
		newDir.TreeID = treeID + "-" + strconv.Itoa(int(newDir.ID))
		return fileDao.Update(&newDir)
	})

	return newDir, errorUtils.Combine(err, FolderCreateError)

}

func (fileService *FileService) MakeAllDir(parentID uint, path string) (lastFile models.File, err error) {
	folderNames := strings.Split(path, "/")
	err = fileService.transaction(func() error {
		for _, folderName := range folderNames {
			if folderName == "" {
				continue
			}

			if lastFile, err = fileService.MakeDir(parentID, folderName); err != nil {
				return errors.Wrapf(err, "创建文件夹%s失败", folderName)
			}
			parentID = lastFile.ID
		}
		return nil
	})
	return lastFile, err
}

func (fileService *FileService) GetAllParentFileList(ID uint) ([]models.File, error) {
	fileList, err := fileService.fileDao.ListAllParents(ID)
	return fileList, err
}

func (fileService *FileService) GetFileInfo(ID uint) (models.File, error) {
	file, err := fileService.fileDao.Find(ID)
	return file, errorUtils.Combine(err, FileNotExitError)
}

func (fileService *FileService) SaveFile(parentID uint, file *multipart.FileHeader, fullPath string) error {
	//将文件保存到本地
	fileName, dst, ext, outerErr := fileUtils.SaveFile(file)
	if outerErr != nil {
		return errors.Wrapf(outerErr, "文件保存失败")
	}

	saveFile := func() error {
		var err error
		var parentFile models.File
		//查看是否需要创建文件夹，并且重置获取父文件夹信息
		if parentFile, err = fileService.resolveParent(parentID, file.Filename, fullPath); err != nil {
			return err
		}

		//本地文件映射
		localFileInfo := models.LocalFile{
			LocalPath: dst,
		}

		if err = fileService.localFileDao.Insert(&localFileInfo); err != nil {
			return err
		}

		oldFile, err := fileService.fileDao.FindByParentIDAndName(parentFile.ParentID, file.Filename)
		if err != nil {
			return errors.Wrapf(err, "文件保存失败")
		}

		if oldFile.ID != 0 {
			//更新
			oldFile.LocalFileID = localFileInfo.ID
			err = fileService.fileDao.Update(&oldFile)
		} else {
			//创建
			fileInfo := models.File{
				Name:        fileName,
				ParentID:    parentFile.ParentID,
				IsDir:       false,
				LocalFileID: localFileInfo.ID,
				ExtType:     ext,
			}

			if err = fileService.fileDao.Insert(&fileInfo); err != nil {
				return errors.Wrapf(err, "文件保存失败")
			}

			//计算新的TreeId

			fileInfo.TreeID = parentFile.TreeID + "-" + strconv.Itoa(int(fileInfo.ID))

			err = fileService.fileDao.Update(&fileInfo)
		}

		return err
	}

	return errors.Wrapf(fileService.transaction(saveFile), "文件保存失败")
}

func (fileService *FileService) resolveParent(parentID models.PrimaryKey, fileName, fullPath string) (models.File, error) {
	parentFile := models.RootFile
	if parentID != 0 {
		parentInfo, err := fileService.fileDao.Find(parentID)
		if err != nil || parentInfo.ID == 0 || parentInfo.IsDir == false {
			return models.File{}, errorUtils.Combine(err, ParentFileNotExitError)
		}
		parentFile = parentInfo
	}

	if fullPath != "" {
		start := 0
		if strings.Index(fullPath, "/") == 0 {
			start = 1
		}
		fullPath = fullPath[start:strings.LastIndex(fullPath, "/"+fileName)]
		if fullPath != "" {
			dir, err := fileService.MakeAllDir(parentFile.ID, fullPath)
			if err != nil {
				return models.File{}, err
			}
			parentFile = dir
		}
	}
	return parentFile, nil
}

func (fileService *FileService) RenameFile(ID uint, newFileName string) error {
	err := fileService.transaction(func() error {
		fileInfo, err := fileService.fileDao.Find(ID)

		if err != nil {
			return errorUtils.Combine(err, FileNotExitError)
		}

		fileInfo.Name = newFileName
		return fileService.fileDao.Update(&fileInfo)
	})
	return errors.Wrapf(err, "重命名失败")
}

func (fileService *FileService) DeleteFile(ID uint) error {
	return errors.Wrapf(fileService.fileDao.Delete(ID), "文件删除失败")
}

func (fileService *FileService) GetFileLocalPathAndFileName(ID uint) (path string, fileName string, err error) {
	var file models.File

	if file, err = fileService.fileDao.Find(ID); err != nil {
		fileName = file.Name
		err = errorUtils.Combine(err, FileNotExitError)
		return
	}

	if file.IsDir {
		fileName = file.Name + ".zip"
		path, err = fileService.toZipFile(file)
		err = errors.Wrapf(err, "zip文件创建失败")
		return
	}

	localFile, err := fileService.localFileDao.Find(file.LocalFileID)

	path = localFile.LocalPath

	err = errors.Wrapf(err, "文件实体丢失")
	return
}

func (fileService *FileService) toZipFile(file models.File) (path string, err error) {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")

	tempPath := "./temp/" + id

	if err = os.MkdirAll(tempPath, os.ModePerm); err != nil {
		return
	}

	files, err := fileService.fileDao.ListByTreeID(file.TreeID)

	if err != nil {
		return
	}

	files = append(files, file)

	root := buildTree(file.ID, files)

	err = fileService.mkDirAndCopyFile(tempPath, root)
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

func (fileService *FileService) mkDirAndCopyFile(bastPath string, root *TreeNode) error {
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
				err := fileService.mkDirAndCopyFile(currentPath, child)
				if err != nil {
					return err
				}
			}
		}
	} else {
		//如果是文件则复制
		localFileInfo, err := fileService.localFileDao.Find(root.FileInfo.LocalFileID)
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
