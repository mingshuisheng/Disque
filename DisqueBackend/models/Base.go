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
