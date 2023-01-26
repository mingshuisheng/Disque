package dao

import (
	"context"
	"disqueBackend/models"
	"disqueBackend/utils/transactionUtils"
	"gorm.io/gorm"
)

type _BaseDao[T any] struct {
	transactionHolder *transactionUtils.TransactionHolder
}

func createBaseDao[T any](transactionHolder *transactionUtils.TransactionHolder) *_BaseDao[T] {
	return &_BaseDao[T]{
		transactionHolder: transactionHolder,
	}
}

func (baseDao *_BaseDao[T]) Insert(model *T) error {
	return baseDao.resolveDB().Create(&model).Error
}

func (baseDao *_BaseDao[T]) Delete(ID models.PrimaryKey) error {
	var model T
	return baseDao.resolveDB().Delete(&model, ID).Error
}

func (baseDao *_BaseDao[T]) Update(model *T) error {
	return baseDao.resolveDB().Save(&model).Error
}

func (baseDao *_BaseDao[T]) Find(ID models.PrimaryKey) (T, error) {
	var model T
	tx := baseDao.resolveDB().First(&model, ID)
	return model, tx.Error
}

func (baseDao *_BaseDao[T]) resolveDB() *gorm.DB {
	return baseDao.transactionHolder.GetDB()
}

func GetDBWithContext(ctx context.Context) *gorm.DB {
	return _DB.WithContext(ctx)
}
