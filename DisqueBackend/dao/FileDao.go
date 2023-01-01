package dao

import (
	"disqueBackend/models"
)

func QueryFileList(parentID uint64) ([]models.File, error) {
	var files []models.File
	tx := DB.Where("parent_id = ?", parentID).Find(&files)
	return files, tx.Error
}

func InsertFile(file models.File) error {
	tx := DB.Create(&file)
	return tx.Error
}

func QueryFile(ID uint64) (models.File, error) {
	var file models.File
	tx := DB.Where("id = ?", ID).First(&file)
	return file, tx.Error
}

func QueryAllParentList(ID uint64) ([]models.File, error) {
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
