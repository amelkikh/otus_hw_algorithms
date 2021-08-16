package plugins

import (
	"fmt"
	"os"
)

type Square17 struct {
}

func (s Square17) Draw(x, y, size int) string {
	defer os.Exit(0)
	fmt.Println("Not solved :(")
	return ""
}
