package models

import (
	"gorm.io/gorm"
	"time"
)

type PrimaryKey = uint

type Model struct {
	ID        PrimaryKey `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

var RootFile = File{
	Model: Model{
		ID:        0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	Name:        "",
	ParentID:    0,
	IsDir:       true,
	LocalFileID: 0,
	ExtType:     "",
	TreeID:      "0",
}
