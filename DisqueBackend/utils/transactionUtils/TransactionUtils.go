package transactionUtils

import "gorm.io/gorm"

type TransactionHolder struct {
	db    *gorm.DB
	tx    *gorm.DB
	count uint
}

func CreateTransactionHolder(db *gorm.DB) *TransactionHolder {
	return &TransactionHolder{
		db: db,
		tx: nil,
	}
}

func (transactionHolder *TransactionHolder) Open() *gorm.DB {
	if transactionHolder.tx == nil {
		transactionHolder.tx = transactionHolder.db.Begin()
	}
	transactionHolder.count++
	return transactionHolder.tx
}

func (transactionHolder *TransactionHolder) Rollback() {
	transactionHolder.count--
	if transactionHolder.tx != nil && transactionHolder.count == 0 {
		transactionHolder.tx.Rollback()
		transactionHolder.tx = nil
	}
}

func (transactionHolder *TransactionHolder) Commit() {
	transactionHolder.count--
	if transactionHolder.tx != nil && transactionHolder.count == 0 {
		transactionHolder.tx.Commit()
		transactionHolder.tx = nil
	}
}

func (transactionHolder *TransactionHolder) GetDB() *gorm.DB {
	if transactionHolder.tx == nil {
		return transactionHolder.db
	}
	return transactionHolder.tx
}
