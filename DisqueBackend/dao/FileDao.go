package dao

import (
	"disqueBackend/models"
	"disqueBackend/utils/transactionUtils"
	"strconv"
	"strings"
)

type FileDao struct {
	*_BaseDao[models.File]
}

func CreateFileDao(holder *transactionUtils.TransactionHolder) *FileDao {
	return &FileDao{
		_BaseDao: createBaseDao[models.File](holder),
	}
}

func (fileDao *FileDao) ListChildren(ID models.PrimaryKey) ([]models.File, error) {
	var files []models.File
	tx := fileDao.resolveDB().Where("parent_id = ?", ID).Find(&files)
	return files, tx.Error
}

func (fileDao *FileDao) FindByParentIDAndName(parentID models.PrimaryKey, fileName string) (models.File, error) {
	var file models.File
	tx := fileDao.resolveDB().Where("parent_id = ?", parentID).Where("name = ?", fileName).Find(&file)
	return file, tx.Error
}

func (fileDao *FileDao) ListByTreeID(treeID string) ([]models.File, error) {
	var files []models.File
	tx := fileDao.resolveDB().Where("tree_id like ?", treeID+"-%").Find(&files)
	return files, tx.Error
}

func (fileDao *FileDao) ListAllParents(ID models.PrimaryKey) ([]models.File, error) {
	db := fileDao.resolveDB()
	var files []models.File
	file, err := fileDao.Find(ID)
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
