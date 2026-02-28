package models

import "gorm.io/gorm"

type Totp struct {
	gorm.Model
	Key     string `gorm:"uniqueIndex"`
	GroupID uint
}
