package window

type index struct{}

func (b *index) Hide() {
	Window.Hide()
}

var Bind = &index{}
