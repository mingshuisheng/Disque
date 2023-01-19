package models

type File struct {
	Model
	Name        string
	ParentID    uint `gorm:"index"`
	IsDir       bool
	LocalFileID uint `json:"-"`
	ExtType     string
	TreeID      string `gorm:"index"`
}

type LocalFile struct {
	Model
	LocalPath string
}
