package window

import "fmt"

type index struct{}

func (b *index) PrintLn(str string) {
	fmt.Println(str)
}

func (b *index) Hide() {
	Window.Hide()
}

var Bind = &index{}
