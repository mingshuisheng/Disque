package dao

import (
	"disqueBackend/models"
	"disqueBackend/utils/transactionUtils"
)

type LocalFileDao struct {
	*_BaseDao[models.LocalFile]
}

func CreateLocalFileDao(holder *transactionUtils.TransactionHolder) *LocalFileDao {
	return &LocalFileDao{
		_BaseDao: createBaseDao[models.LocalFile](holder),
	}
}
