package service

import (
	"context"
	"disqueBackend/dao"
	"disqueBackend/utils/transactionUtils"
)

type _BaseService struct {
	ctx context.Context
	//db  *gorm.DB
	transactionHolder *transactionUtils.TransactionHolder
}

func createBaseService(ctx context.Context) *_BaseService {
	return &_BaseService{
		ctx:               ctx,
		transactionHolder: transactionUtils.CreateTransactionHolder(dao.GetDBWithContext(ctx)),
	}
}

func (baseService _BaseService) transaction(fc func() (err error)) error {

	defer func() {
		if rec := recover(); rec != nil {
			baseService.transactionHolder.Rollback()
		}
	}()

	baseService.transactionHolder.Open()
	if err := fc(); err != nil {
		baseService.transactionHolder.Rollback()
		return err
	}
	baseService.transactionHolder.Commit()
	return nil
}
