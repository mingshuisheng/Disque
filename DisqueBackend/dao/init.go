package dao

import (
	"disqueBackend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _DB *gorm.DB

var FileDao *_FileDao
var LocalFileDao *_LocalFileDao

func Init() {
	var db *gorm.DB
	var err error

	if db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{}); err != nil {
		panic("failed to connect database")
	}

	if err = db.AutoMigrate(&models.File{}); err != nil {
		panic("failed to migrate models.File")
	}

	if err = db.AutoMigrate(&models.LocalFile{}); err != nil {
		panic("failed to migrate models.LocalFile")
	}

	_DB = db
	FileDao = &_FileDao{_BaseDao: new(_BaseDao[models.File])}
	LocalFileDao = &_LocalFileDao{_BaseDao: new(_BaseDao[models.LocalFile])}
}
