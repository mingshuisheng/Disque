package dao

import (
	"disqueBackend/models"
	"disqueBackend/utils/transactionUtils"
)

type FileCacheDao struct {
	*_BaseDao[models.File]
}

func CreateFileCacheDao(holder *transactionUtils.TransactionHolder) *FileCacheDao {
	return &FileCacheDao{
		_BaseDao: createBaseDao[models.File](holder),
	}
}
