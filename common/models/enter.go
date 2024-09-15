package models

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
}

type PageInfo struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
