package dao

import (
	"disqueBackend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _DB *gorm.DB

func Init() {
	if db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}); err != nil {
		panic("failed to connect database")
	} else {
		_DB = db
	}

	if err := _DB.AutoMigrate(&models.File{}); err != nil {
		panic("failed to migrate models.File")
	}

	if err := _DB.AutoMigrate(&models.FileCache{}); err != nil {
		panic("failed to migrate models.FileCache")
	}

	if err := _DB.AutoMigrate(&models.LocalFile{}); err != nil {
		panic("failed to migrate models.LocalFile")
	}

}
