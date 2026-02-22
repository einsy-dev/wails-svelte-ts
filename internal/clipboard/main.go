package clipboard

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/einsy-dev/WailsSvelte/utils"
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
	go b.Start()
}

func (b *cp) Start() {
	tCh, iCh := Watch(100 * time.Microsecond)

	for {
		select {
		case msg := <-tCh:
			b.text = utils.FilterUnique(append([]string{string(msg)}, b.text...))
		case msg := <-iCh:
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
