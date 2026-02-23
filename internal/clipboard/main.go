package clipboard

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/einsy-dev/WailsSvelte/internal/app"
	m "github.com/einsy-dev/WailsSvelte/internal/storage/models"
	"golang.design/x/clipboard"
)

type cp struct {
	text []string
	img  []string
}

func (b *cp) Startup() {
	clipboard.Init()
	go b.Start()
}

func (b *cp) Start() {
	tCh, iCh := Watch(100 * time.Microsecond)
	var hGroup m.Group
	app.Db.Where(&m.Group{Name: "History"}).First(&hGroup)

	for {
		select {
		case msg := <-tCh:
			app.Db.Create(&m.Text{Value: string(msg), GroupID: hGroup.ID})
		case msg := <-iCh:
			b.img = append(b.img, string(msg))
			home, _ := os.UserHomeDir()
			path := filepath.Join(home, "Desktop", "clipboard_image.png")
			os.WriteFile(path, msg, 0644)
		case <-app.Ctx.Done():
			fmt.Println("Stopping watcher...")
			return
		}
	}

}

var Clipboard = &cp{}
