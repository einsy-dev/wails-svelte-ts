package window

import (
	"context"

	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type window struct {
	ctx context.Context
	Pid int
}

func (b *window) Startup(ctx context.Context) {
	b.ctx = ctx
}

func (b *window) Show() {
	b.Pid = robotgo.GetPid()

	windowW, windowH := runtime.WindowGetSize(b.ctx) // window size
	screenW, screenH := robotgo.GetScreenSize()      // screen size
	x, y := robotgo.Location()                       //cursor position

	newX, newY := x, y

	if x > screenW-windowW {
		newX = x - windowW
	}

	if y > screenH-windowH {
		newY = y - windowH
	}

	runtime.WindowSetPosition(b.ctx, newX, newY)
	runtime.Show(b.ctx)
}

func (b *window) Hide() {
	runtime.WindowHide(b.ctx)
}

var Window = &window{}
