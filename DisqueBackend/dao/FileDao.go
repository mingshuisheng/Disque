package dao

import (
	"disqueBackend/models"
)

func QueryFileList(parentID uint) ([]models.File, error) {
	var files []models.File
	tx := DB.Where("parent_id = ?", parentID).Find(&files)
	return files, tx.Error
}

func InsertFile(file models.File) error {
	tx := DB.Create(&file)
	return tx.Error
}

func QueryFile(ID uint) (models.File, error) {
	var file models.File
	tx := DB.Where("id = ?", ID).First(&file)
	return file, tx.Error
}

func QueryAllParentList(ID uint) ([]models.File, error) {
	var files []models.File
	file, err := QueryFile(ID)
	if err != nil {
		return files, err
	}
	files = append(files, file)
	for file.ParentID != 0 {
		file, err = QueryFile(file.ParentID)
		if err != nil {
			return files, err
		}
	}

	return files, nil
}

func InsertLocalFile(localFile *models.LocalFile) error {
	tx := DB.Create(localFile)
	return tx.Error
}

func QueryLocalFile(ID uint) (models.LocalFile, error) {
	localFile := models.LocalFile{}
	tx := DB.First(&localFile, ID)
	return localFile, tx.Error
}
