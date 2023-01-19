package dao

import (
	"disqueBackend/models"
	"strconv"
	"strings"
)

func QueryFileList(parentID uint) ([]models.File, error) {
	var files []models.File
	tx := DB.Where("parent_id = ?", parentID).Find(&files)
	return files, tx.Error
}

func QueryFileByParentIDAndName(parentID uint, fileName string) (models.File, error) {
	var file models.File
	tx := DB.Where("parent_id = ?", parentID).Where("name = ?", fileName).Find(&file)
	return file, tx.Error
}

func QueryAllChildrenByTreeID(treeID string) ([]models.File, error) {
	var files []models.File
	tx := DB.Where("tree_id like ?", treeID+"-%").Find(&files)
	return files, tx.Error
}

func InsertFile(file *models.File) error {
	tx := DB.Create(file)
	return tx.Error
}

func UpdateFile(file *models.File) error {
	tx := DB.Save(file)
	return tx.Error
}

func DeleteFile(ID uint) error {
	tx := DB.Delete(&models.File{}, ID)
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

	parents := strings.Split(file.TreeID, "-")

	var parentIDList []uint

	parentIDList = append(parentIDList, file.ID)

	for _, str := range parents {
		if str != "0" {
			id, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				return files, err
			}
			parentIDList = append(parentIDList, uint(id))
		}
	}

	tx := DB.Where("ID IN ?", parentIDList).Find(&files)
	if tx.Error != nil {
		return files, tx.Error
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
