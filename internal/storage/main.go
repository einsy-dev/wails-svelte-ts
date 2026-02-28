package storage

import (
	"github.com/einsy-dev/WailsSvelte/internal/app"
	m "github.com/einsy-dev/WailsSvelte/internal/storage/models"
	"gorm.io/gorm/clause"
)

func Seed() {
	var groups = []m.Group{{Name: "history"}, {Name: "favorites"}, {Name: "passwords"}, {Name: "totp"}}
	app.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&groups)
}
