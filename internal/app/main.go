package app

import (
	"context"
)

type app struct {
	ctx context.Context
}

func (b *app) Startup(ctx context.Context) {
	b.ctx = ctx
	// time.Sleep(8 * time.Second)
	// runtime.WindowSetPosition(ctx, 1400, 700)
	// runtime.Show(ctx)
}

var App = &app{}
