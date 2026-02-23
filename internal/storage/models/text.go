package models

import "gorm.io/gorm"

type Text struct {
	gorm.Model
	Value   string
	GroupID uint
}
