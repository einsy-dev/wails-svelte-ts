package window

import (
	"context"
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type window struct {
	ctx context.Context
}

func (b *window) Startup(ctx context.Context) {
	b.ctx = ctx
}

func (b *window) Show() {
	windowW, windowH := runtime.WindowGetSize(b.ctx) // window size
	screenW, screenH := robotgo.GetScreenSize()      // screen size
	x, y := robotgo.Location()                       //cursor position

	runtime.WindowSetPosition(b.ctx, x, y)
	fmt.Println(windowW, windowH, screenW, screenH)
}

var Window = &window{}
