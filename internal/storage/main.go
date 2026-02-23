package storage

import (
	"github.com/einsy-dev/WailsSvelte/internal/app"
	m "github.com/einsy-dev/WailsSvelte/internal/storage/models"
	"gorm.io/gorm/clause"
)

func Seed() {
	var groups = []m.Group{{Name: "History"}, {Name: "Favorite"}}
	app.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&groups)

	var g m.Group
	app.Db.Where(&m.Group{Name: "Favorite"}).First(&g)
}
