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
	"gorm.io/gorm"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

type FileService struct {
	BaseService
}

func CreateFileService(ctx context.Context) *FileService {
	return &FileService{
		BaseService{ctx, nil},
	}
}

func (fileService *FileService) GetFileList(ID models.PrimaryKey) ([]models.File, error) {
	fileList, err := dao.FileDao.ListChildren(ID, fileService.getDB())
	return fileList, errors.Wrapf(err, "文件不存在")
}

func (fileService *FileService) makeDir(db *gorm.DB, parentID models.PrimaryKey, name string) (newDir models.File, err error) {
	treeID := models.RootFile.TreeID
	if parentID != 0 {
		info, err := dao.FileDao.Find(parentID, db)
		if err != nil {
			return models.File{}, errorUtils.Combine(err, ParentFileNotExitError)
		}
		treeID = info.TreeID
	}

	newDir = models.File{ParentID: parentID, Name: name, IsDir: true}

	oldFile, err := dao.FileDao.FindByParentIDAndName(parentID, name, db)

	if err != nil {
		err = errorUtils.Combine(err, ParentFileNotExitError)
		return
	}

	if oldFile.ID != 0 {
		return newDir, nil
	}

	err = dao.FileDao.Insert(&newDir, db)
	if err != nil {
		return newDir, errorUtils.Combine(err, FolderCreateError)
	}
	newDir.TreeID = treeID + "-" + strconv.Itoa(int(newDir.ID))
	err = dao.FileDao.Update(&newDir, db)

	return newDir, errorUtils.Combine(err, FolderCreateError)

}

func (fileService *FileService) MakeDir(parentID uint, name string) (models.File, error) {
	var newDir models.File
	err := fileService.transaction(func(tx *gorm.DB) error {
		resultDir, innErr := fileService.makeDir(tx, parentID, name)
		newDir = resultDir
		return innErr
	})
	return newDir, err
}

func (fileService *FileService) makeAllDir(db *gorm.DB, parentID uint, path string) (lastFile models.File, err error) {
	folderNames := strings.Split(path, "/")

	for _, folderName := range folderNames {
		if folderName == "" {
			continue
		}

		if lastFile, err = fileService.makeDir(db, parentID, folderName); err != nil {
			return models.File{}, errors.Wrapf(err, "创建文件夹%s失败", folderName)
		}
		parentID = lastFile.ID
	}
	return lastFile, nil
}

func (fileService *FileService) MakeAllDir(parentID uint, path string) (models.File, error) {
	var lastFile models.File
	err := fileService.transaction(func(tx *gorm.DB) error {
		dir, err := fileService.makeAllDir(tx, parentID, path)
		lastFile = dir
		return err
	})
	return lastFile, err
}

func (fileService *FileService) GetAllParentFileList(ID uint) ([]models.File, error) {
	fileList, err := dao.FileDao.ListAllParents(ID, fileService.getDB())
	return fileList, err
}

func (fileService *FileService) GetFileInfo(ID uint) (models.File, error) {
	file, err := dao.FileDao.Find(ID, fileService.getDB())
	return file, errorUtils.Combine(err, FileNotExitError)
}

func (fileService *FileService) SaveFile(parentID uint, file *multipart.FileHeader, fullPath string) error {

	//将文件保存到本地
	fileName, dst, ext, outterErr := fileUtils.SaveFile(file)
	if outterErr != nil {
		return errors.Wrapf(outterErr, "文件保存失败")
	}

	outterErr = fileService.transaction(func(tx *gorm.DB) error {
		var err error
		var parentFile models.File
		//查看是否需要创建文件夹，并且重置获取父文件夹信息
		if parentFile, err = fileService.resolveParent(tx, parentID, file.Filename, fullPath); err != nil {
			return err
		}

		//本地文件映射
		localFileInfo := models.LocalFile{
			LocalPath: dst,
		}

		if err = dao.LocalFileDao.Insert(&localFileInfo, tx); err != nil {
			return err
		}

		oldFile, err := dao.FileDao.FindByParentIDAndName(parentFile.ParentID, file.Filename, tx)
		if err != nil {
			return errors.Wrapf(err, "文件保存失败")
		}

		if oldFile.ID != 0 {
			//更新
			oldFile.LocalFileID = localFileInfo.ID
			err = dao.FileDao.Update(&oldFile, tx)
		} else {
			//创建
			fileInfo := models.File{
				Name:        fileName,
				ParentID:    parentFile.ParentID,
				IsDir:       false,
				LocalFileID: localFileInfo.ID,
				ExtType:     ext,
			}

			if err = dao.FileDao.Insert(&fileInfo, tx); err != nil {
				return errors.Wrapf(err, "文件保存失败")
			}

			//计算新的TreeId

			fileInfo.TreeID = parentFile.TreeID + "-" + strconv.Itoa(int(fileInfo.ID))

			err = dao.FileDao.Update(&fileInfo, tx)
		}

		return err
	})

	return errors.Wrapf(outterErr, "文件保存失败")
}

func (fileService *FileService) resolveParent(db *gorm.DB, parentID models.PrimaryKey, fileName, fullPath string) (models.File, error) {
	parentFile := models.RootFile
	if parentID != 0 {
		parentInfo, err := dao.FileDao.Find(parentID, db)
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
			dir, err := fileService.makeAllDir(db, parentFile.ID, fullPath)
			if err != nil {
				return models.File{}, err
			}
			parentFile = dir
		}
	}
	return parentFile, nil
}

func (fileService *FileService) RenameFile(ID uint, newFileName string) error {
	err := fileService.transaction(func(tx *gorm.DB) error {
		fileInfo, err := dao.FileDao.Find(ID, tx)

		if err != nil {
			return errorUtils.Combine(err, FileNotExitError)
		}

		fileInfo.Name = newFileName
		return dao.FileDao.Update(&fileInfo, tx)
	})
	return errors.Wrapf(err, "重命名失败")
}

func (fileService *FileService) DeleteFile(ID uint) error {
	return errors.Wrapf(dao.FileDao.Delete(ID, fileService.getDB()), "文件删除失败")
}

func (fileService *FileService) GetFileLocalPathAndFileName(ID uint) (path string, fileName string, err error) {
	db := fileService.getDB()
	var file models.File

	if file, err = dao.FileDao.Find(ID, db); err != nil {
		fileName = file.Name
		err = errorUtils.Combine(err, FileNotExitError)
		return
	}

	if file.IsDir {
		fileName = file.Name + ".zip"
		path, err = fileService.toZipFile(db, file)
		err = errors.Wrapf(err, "zip文件创建失败")
		return
	}

	localFile, err := dao.LocalFileDao.Find(file.LocalFileID, db)

	path = localFile.LocalPath

	err = errors.Wrapf(err, "文件实体丢失")
	return
}

func (fileService *FileService) toZipFile(db *gorm.DB, file models.File) (path string, err error) {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")

	tempPath := "./temp/" + id

	if err = os.MkdirAll(tempPath, os.ModePerm); err != nil {
		return
	}

	files, err := dao.FileDao.ListByTreeID(file.TreeID, db)

	if err != nil {
		return
	}

	files = append(files, file)

	root := buildTree(file.ID, files)

	err = fileService.mkDirAndCopyFile(db, tempPath, root)
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

func (fileService *FileService) mkDirAndCopyFile(db *gorm.DB, bastPath string, root *TreeNode) error {
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
				err := fileService.mkDirAndCopyFile(db, currentPath, child)
				if err != nil {
					return err
				}
			}
		}
	} else {
		//如果是文件则复制
		localFileInfo, err := dao.LocalFileDao.Find(root.FileInfo.LocalFileID, db)
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
