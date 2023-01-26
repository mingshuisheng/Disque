package dao

import (
	"disqueBackend/models"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type _FileDao struct {
	*_BaseDao[models.File]
}

func (*_FileDao) ListChildren(ID models.PrimaryKey, dbs *gorm.DB) ([]models.File, error) {
	var files []models.File
	tx := resolveDB(dbs).Where("parent_id = ?", ID).Find(&files)
	return files, tx.Error
}

func (*_FileDao) FindByParentIDAndName(parentID models.PrimaryKey, fileName string, dbs *gorm.DB) (models.File, error) {
	var file models.File
	tx := resolveDB(dbs).Where("parent_id = ?", parentID).Where("name = ?", fileName).Find(&file)
	return file, tx.Error
}

func (*_FileDao) ListByTreeID(treeID string, dbs *gorm.DB) ([]models.File, error) {
	var files []models.File
	tx := resolveDB(dbs).Where("tree_id like ?", treeID+"-%").Find(&files)
	return files, tx.Error
}

func (dao *_FileDao) ListAllParents(ID models.PrimaryKey, dbs *gorm.DB) ([]models.File, error) {
	db := resolveDB(dbs)
	var files []models.File
	file, err := dao.Find(ID, db)
	if err != nil {
		return files, err
	}

	parents := strings.Split(file.TreeID, "-")

	var parentIDList []models.PrimaryKey

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

	tx := db.Where("ID IN ?", parentIDList).Find(&files)
	return files, tx.Error
}
