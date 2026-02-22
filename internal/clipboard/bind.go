package clipboard

import (
	"github.com/einsy-dev/WailsSvelte/internal/window"
	"github.com/go-vgo/robotgo"
	"golang.design/x/clipboard"
)

type index struct {
}

func (b *index) GetHistory() []string {
	return Clipboard.text
}

func (b *index) GetGroup(name string) []string {
	return Clipboard.text
}

func (b *index) Paste(val string) {
	robotgo.ActivePid(window.Window.Pid)
	clipboard.Write(clipboard.FmtText, []byte(val))
	robotgo.KeyTap("v", "ctrl")
}

var Bind = &index{}
