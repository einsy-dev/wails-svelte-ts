package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex"`
	Logo  *string
	Texts []Text
	Imgs  []Img
	Totp  []Totp
}
