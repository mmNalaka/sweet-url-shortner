package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
}

func (m *Tag) TableName() string {
    return "tag"
}
