package model

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	Id            uint   `gorm:"primarykey"`
	Domain        string `gorm:"not null"`
	Key           string `gorm:"not null"`
	Url           string `gorm:"not null"`
	Proxy         bool   `gorm:"default:false"`
	Title         string
	Description   string
	Image         string
	Password      string
	Rewrite       bool      `gorm:"default:false"`
	Archived      bool      `gorm:"default:false"`
	Clicks        uint      `gorm:"default:0"`
	ExpiresAt     time.Time `gorm:"default:null"`
	ExpirationUrl string    `gorm:"default:null"`
	Target        string
	Type          string `gorm:"default:'redirect'"`
	LastClicked   time.Time
	UserId        uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (m *Link) TableName() string {
	return "link"
}
