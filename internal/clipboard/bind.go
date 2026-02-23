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
	app.Db.Where(&m.Group{Name: "History"}).First(&hGroup)

	var res []m.Text
	app.Db.Where("group_id = ? AND value <> ? AND value IS NOT NULL", hGroup.ID, "").Limit(-1).Order("id desc").Find(&res)
	fmt.Println(len(res))
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

var Bind = &index{}
