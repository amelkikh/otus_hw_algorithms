package square

import (
	"fmt"
)

type IDrawFunc interface {
	Draw(x, y, size int) string
}

type IDraw interface {
	Draw()
}

type ISquare interface {
	IDraw
	SetDrawFunc(f IDrawFunc)
}

type square struct {
	size     int
	drawFunc IDrawFunc
}

func (s square) Draw() {
	for y := 0; y < s.size; y++ {
		for x := 0; x < s.size; x++ {
			fmt.Print(s.drawFunc.Draw(x, y, s.size))
		}
		fmt.Println()
	}
}

func NewSquare(size int, f IDrawFunc) ISquare {
	return &square{
		size:     size,
		drawFunc: f,
	}
}

func (s *square) SetDrawFunc(f IDrawFunc) {
	s.drawFunc = f
}
