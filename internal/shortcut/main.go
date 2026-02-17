package shortcut

import (
	"fmt"
	"log"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func Startup() {
	fmt.Println("shortcut start")
	mainthread.Init(fn)
}

func fn() {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModWin}, hotkey.KeyV)
	err := hk.Register()

	if err != nil {
		log.Fatalf("hotkey: failed to register hotkey: %v", err)
		return
	}

	for msg := range hk.Keydown() {

		log.Printf("hotkey: %v is down\n", msg)
	}
}
