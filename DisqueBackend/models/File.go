package models

import "time"

type File struct {
	Model
	Name        string
	ParentID    uint `gorm:"index"`
	IsDir       bool
	LocalFileID uint `json:"-"`
	ExtType     string
	TreeID      string `gorm:"index"`
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
