package logic

import (
	"disqueBackend/dao"
	"disqueBackend/models"
	"errors"
)

func GetFileList(parentID uint) []models.File {
	fileList, err := dao.QueryFileList(parentID)
	if err != nil {
		return []models.File{}
	}
	return fileList
}

func MakeDir(parentID uint, name string) error {
	newDir := models.File{ParentID: parentID, Name: name, IsDir: true}
	err := dao.InsertFile(newDir)
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

func SaveFileInfo(parentID uint, fileName string, dst string, ext string) (err error) {
	//校验parentID是否有效
	if parentID != 0 {
		info, err := GetFileInfo(parentID)
		if err != nil || info.ID == 0 || info.IsDir == false {
			err = errors.New("正在操作的文件不存在")
			return err
		}
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

	err = dao.InsertFile(fileInfo)

	return err
}

func GetFileLocalPathAndFileName(ID uint) (path string, fileName string, err error) {
	path = ""
	fileName = ""
	file, err := dao.QueryFile(ID)
	if err != nil {
		return
	}

	localFile, err := dao.QueryLocalFile(file.LocalFileID)
	if err != nil {
		return
	}

	path = localFile.LocalPath
	fileName = file.Name
	return
}
