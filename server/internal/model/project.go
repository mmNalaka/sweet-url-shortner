package model

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	Id         uint   `gorm:"primarykey"`
	Name       string `gorm:"not null"`
	Slug       string `gorm:"unique;not null"`
	Logo       string
	InviteCode string `gorm:"unique"`
	OwnerId    uint
	Owner      User
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (m *Project) TableName() string {
	return "project"
}
