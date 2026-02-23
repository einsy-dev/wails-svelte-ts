package main

import (
	"context"
	"embed"

	"github.com/einsy-dev/WailsSvelte/internal/app"
	"github.com/einsy-dev/WailsSvelte/internal/clipboard"
	"github.com/einsy-dev/WailsSvelte/internal/shortcut"
	"github.com/einsy-dev/WailsSvelte/internal/storage"
	"github.com/einsy-dev/WailsSvelte/internal/window"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	err := wails.Run(&options.App{
		Title:             "WailsSvelte",
		Width:             500,
		Height:            300,
		DisableResize:     true,
		Frameless:         true,
		StartHidden:       true,
		HideWindowOnClose: true,
		AlwaysOnTop:       true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup: func(ctx context.Context) {
			app.Startup(ctx)
			shortcut.Startup()
			clipboard.Clipboard.Startup()
			storage.Seed()
		},
		Bind: []interface{}{
			clipboard.Bind,
			window.Bind,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
