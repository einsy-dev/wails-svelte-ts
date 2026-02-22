package storage

import (
	"context"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Startup(ctx context.Context) {
	db, err := gorm.Open(sqlite.Open("storage.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Text{}, &Img{}, &Totp{}, &Group{})
	Storage.db = db
}

type storage struct {
	ctx context.Context
	db  *gorm.DB
}

var Storage = &storage{}
