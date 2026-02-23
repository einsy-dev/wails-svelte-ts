package window

import (
	"github.com/einsy-dev/WailsSvelte/internal/app"
	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type window struct {
	Pid int
}

func (b *window) Show() {
	b.Pid = robotgo.GetPid()

	windowW, windowH := runtime.WindowGetSize(app.Ctx) // window size
	screenW, screenH := robotgo.GetScreenSize()        // screen size
	x, y := robotgo.Location()                         //cursor position

	newX, newY := x, y

	if x > screenW-windowW {
		newX = x - windowW
	}

	if y > screenH-windowH {
		newY = y - windowH
	}

	runtime.WindowSetPosition(app.Ctx, newX, newY)
	runtime.Show(app.Ctx)
}

func (b *window) Hide() {
	runtime.WindowHide(app.Ctx)
}

var Window = &window{}
