package services

import (
	"fmt"

	"github.com/einsy-dev/WailsSvelte/internal/app"
	m "github.com/einsy-dev/WailsSvelte/internal/storage/models"
)

type text struct{}

func (t *text) Create(text *m.Text) {
	app.Db.FirstOrCreate(&text)
}
func (t *text) Read(groupID uint) []m.Text {
	var text []m.Text
	err := app.Db.Where("group_id = ?", groupID).Find(&text).Error
	if err != nil {
		fmt.Println(err)
	}
	return text
}

func (t *text) Update() {}
func (t *text) Delete() {}

var Text = &text{}
