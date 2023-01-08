package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name        string
	ParentID    uint `gorm:"index"`
	IsDir       bool
	LocalFileID uint `json:"-"`
	ExtType     string
	TreeID      string `gorm:"index"`
}

type LocalFile struct {
	gorm.Model
	LocalPath string
}
