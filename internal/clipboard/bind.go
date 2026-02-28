package clipboard

import (
	"fmt"

	"github.com/einsy-dev/WailsSvelte/internal/app"
	m "github.com/einsy-dev/WailsSvelte/internal/storage/models"
	"github.com/einsy-dev/WailsSvelte/internal/window"
	"github.com/go-vgo/robotgo"
	"golang.design/x/clipboard"
)

type index struct{}

func (b *index) GetHistory() []m.Text {

	var hGroup m.Group
	err := app.Db.Where(&m.Group{Name: "history"}).First(&hGroup).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var res []m.Text
	app.Db.Where("group_id = ? AND value <> ? AND value IS NOT NULL", hGroup.ID, "").Limit(-1).Order("id desc").Find(&res)
	return res
}

func (b *index) GetByGroup(name string) []m.Text {
	var res []m.Text
	var hGroup m.Group
	app.Db.Where(&m.Group{Name: name}).First(&hGroup)
	app.Db.Where("group_id = ?", hGroup.ID).Find(&res)
	return res
}

func (b *index) Paste(val string) {
	robotgo.ActivePid(window.Window.Pid)
	clipboard.Write(clipboard.FmtText, []byte(val))
	robotgo.KeyTap("v", "ctrl")
}

func (b *index) SetGroup(id uint, name string) {
	var group m.Group
	err := app.Db.Where(&m.Group{Name: name}).First(&group).Error
	if err != nil {
		fmt.Println("find group err", err)
		return
	}
	var item m.Text
	app.Db.Where("id = ?", id).First(&item)
	item.GroupID = group.ID
	app.Db.Save(&item)
}

var Bind = &index{}
