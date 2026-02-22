package clipboard

import (
	"bytes"
	"time"

	"golang.design/x/clipboard"
)

func Watch(delay time.Duration) (chan string, chan []byte) {
	tCh := make(chan string)
	iCh := make(chan []byte)

	go func() {
		ticker := time.NewTicker(delay)
		var lastText []byte
		var lastImg []byte

		for range ticker.C {
			text := clipboard.Read(clipboard.FmtText)
			if !bytes.Equal(text, lastText) {
				lastText = text
				tCh <- string(text)
			}

			img := clipboard.Read(clipboard.FmtImage)
			if !bytes.Equal(img, lastImg) {
				lastImg = img
				iCh <- img
			}
		}
	}()

	return tCh, iCh
}
