package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func (Category) TableName() string {
	return "category"
}
