package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name      string
	ParentID  uint64 `gorm:"index"`
	IsDir     bool
	LocalPath string `json:"-"`
}
