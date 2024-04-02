package model

import "gorm.io/gorm"

type Link struct {
	gorm.Model
}

func (m *Link) TableName() string {
    return "link"
}
