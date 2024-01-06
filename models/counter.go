package models

import (
	"gorm.io/gorm"
)

type Counter struct {
	gorm.Model
	Id  int64 
	Count int64  `gorm:"default:0"`
}