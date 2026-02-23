package models

import "gorm.io/gorm"

type Totp struct {
	gorm.Model
	Key     string
	GroupID uint
}
