package dao

import (
	"context"
	"disqueBackend/models"
	"gorm.io/gorm"
)

type _BaseDao[T any] int

func (*_BaseDao[T]) Insert(model *T, dbs *gorm.DB) error {
	return resolveDB(dbs).Create(&model).Error
}

func (*_BaseDao[T]) Delete(ID models.PrimaryKey, dbs *gorm.DB) error {
	var model T
	return resolveDB(dbs).Delete(&model, ID).Error
}

func (*_BaseDao[T]) Update(model *T, dbs *gorm.DB) error {
	return resolveDB(dbs).Save(&model).Error
}

func (*_BaseDao[T]) Find(ID models.PrimaryKey, dbs *gorm.DB) (T, error) {
	var model T
	tx := resolveDB(dbs).First(&model, ID)
	return model, tx.Error
}

type TransactionCallback = func(tx *gorm.DB) error

func Transaction(db *gorm.DB, fc TransactionCallback) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		return fc(tx)
	})
}

func GetDBWithContext(ctx context.Context) *gorm.DB {
	return _DB.WithContext(ctx)
}

func resolveDB(dbs ...*gorm.DB) *gorm.DB {
	if len(dbs) > 0 && dbs[0] != nil {
		return dbs[0]
	}
	return _DB.WithContext(context.Background())
}
