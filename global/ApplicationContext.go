package global

import (
	"gorm.io/gorm"
)

type ApplicationContext struct {
	DatabaseConnection *gorm.DB
}

var Ctx *ApplicationContext
