package service

import (
	"context"
	"disqueBackend/dao"
	"gorm.io/gorm"
)

type BaseService struct {
	ctx context.Context
	db  *gorm.DB
}

func (baseService BaseService) transaction(fc func(tx *gorm.DB) (err error)) error {
	return dao.Transaction(baseService.getDB(), fc)
}

func (baseService BaseService) getDB() *gorm.DB {
	if baseService.db == nil {
		baseService.db = dao.GetDBWithContext(baseService.ctx)
	}
	return baseService.db
}
