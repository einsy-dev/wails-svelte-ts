package clipboard

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"golang.design/x/clipboard"
)

type cp struct {
	ctx  context.Context
	text []string
	img  []string
}

func (b *cp) Startup(ctx context.Context) {
	b.ctx = ctx
	clipboard.Init()
	b.Watch()
}

func (b *cp) Watch() {
	fmt.Println("start")
	chs := clipboard.Watch(b.ctx, clipboard.FmtText)
	chi := clipboard.Watch(b.ctx, clipboard.FmtImage)

	for {
		select {
		case msg := <-chs:
			b.text = append(b.text, string(msg))
			fmt.Println(string(msg), b.text)
		case msg := <-chi:
			b.img = append(b.img, string(msg))
			home, _ := os.UserHomeDir()
			path := filepath.Join(home, "Desktop", "clipboard_image.png")
			os.WriteFile(path, msg, 0644)
		case <-b.ctx.Done():
			fmt.Println("Stopping watcher...")
			return
		}
	}

}

var Clipboard = &cp{}
