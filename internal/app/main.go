package app

import (
	"context"

	m "github.com/einsy-dev/WailsSvelte/internal/storage/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// global state

var Ctx context.Context
var Db *gorm.DB

func Startup(ctx context.Context) {
	Ctx = ctx

	db, err := gorm.Open(sqlite.Open("storage.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&m.Text{}, &m.Img{}, &m.Totp{}, &m.Group{})
	Db = db
}
