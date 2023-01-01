package logic

import (
	"disqueBackend/dao"
	"disqueBackend/models"
)

func GetFileList(parentID uint64) []models.File {
	fileList, err := dao.QueryFileList(parentID)
	if err != nil {
		return []models.File{}
	}
	return fileList
}

func MakeDir(parentID uint64, name string) error {
	newDir := models.File{ParentID: parentID, Name: name, IsDir: true}
	err := dao.InsertFile(newDir)
	if err != nil {
		return err
	}
	return nil
}

func GetAllParentFileList(ID uint64) []models.File {
	fileList, err := dao.QueryAllParentList(ID)
	if err != nil {
		return []models.File{}
	}

	return fileList
}

func GetFileInfo(ID uint64) (models.File, error) {
	return dao.QueryFile(ID)
}
