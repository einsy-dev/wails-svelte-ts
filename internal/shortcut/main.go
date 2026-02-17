package shortcut

import (
	"fmt"
	"log"

	"github.com/einsy-dev/WailsSvelte/internal/window"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func Startup() {
	fmt.Println("shortcut start")
	mainthread.Init(fn)
}

func fn() {
	RegistryEdit()

	hk := hotkey.New([]hotkey.Modifier{hotkey.ModWin}, hotkey.KeyV)
	err := hk.Register()

	if err != nil {
		log.Fatalf("hotkey: failed to register hotkey: %v", err)
		return
	}

	for range hk.Keydown() {
		window.Window.Show()
	}
}
