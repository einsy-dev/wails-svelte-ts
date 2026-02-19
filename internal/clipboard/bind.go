package clipboard

import (
	"fmt"
)

type index struct {
}

func (b *index) Test() {
	fmt.Println("clicl")
}

var Bind = &index{}
