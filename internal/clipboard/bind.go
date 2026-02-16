package clipboard

import (
	"context"
	"fmt"
)

type index struct {
	ctx context.Context
}

func (b *index) Test(ctx context.Context) {
	fmt.Println("clicl")
}

var Bind = &index{}
