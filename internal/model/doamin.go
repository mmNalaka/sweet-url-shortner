package model

import "gorm.io/gorm"

type Doamin struct {
	gorm.Model
}

func (m *Doamin) TableName() string {
    return "doamin"
}
