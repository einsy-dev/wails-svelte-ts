package storage

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Title string
	Logo  string
	Texts []Text
}

type Text struct {
	gorm.Model
	Value   string
	GroupId uint
}
