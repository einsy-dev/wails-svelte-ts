package services

import (
	"github.com/einsy-dev/WailsSvelte/internal/app"
	m "github.com/einsy-dev/WailsSvelte/internal/storage/models"
)

type group struct{}

func (t *group) Create(name string) *m.Group {
	var group = &m.Group{Name: name}
	app.Db.FirstOrCreate(&group)
	return group
}

func (t *group) Read(group *m.Group) {
	app.Db.First(group)
}

func (t *group) Update() {}
func (t *group) Delete() {}

var Group = &group{}
