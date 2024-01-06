package models

import "gorm.io/gorm"
type Quote struct {
	gorm.Model
    Greek_quote string 
	English_translation string 
	Description string
}