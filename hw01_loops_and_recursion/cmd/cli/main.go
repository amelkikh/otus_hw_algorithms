package main

import (
	"flag"
	"fmt"

	"github.com/amelkikh/otus_hw_algorithms/hw01_loops_and_recursion/internal/pkg/plugins"
	"github.com/amelkikh/otus_hw_algorithms/hw01_loops_and_recursion/internal/square"
)

func main() {
	squareNumFlag := flag.Uint("id", 1, "Sequence number of square from homework")
	flag.Parse()
	nums := *squareNumFlag
	var f square.IDrawFunc
	squareFuncs := []square.IDrawFunc{
		&plugins.Square1{},
		&plugins.Square2{},
		&plugins.Square3{},
		&plugins.Square4{},
		&plugins.Square5{},
		&plugins.Square6{},
		&plugins.Square7{},
		&plugins.Square8{},
		&plugins.Square9{},
		&plugins.Square10{},
		&plugins.Square11{},
		&plugins.Square12{},
		&plugins.Square13{},
		&plugins.Square14{},
		&plugins.Square15{},
		&plugins.Square16{},
		&plugins.Square17{},
		&plugins.Square18{},
		&plugins.Square19{},
		&plugins.Square20{},
		&plugins.Square21{},
		&plugins.Square22{},
		&plugins.Square23{},
		&plugins.Square24{},
		&plugins.Square25{},
	}
	if int(nums) > len(squareFuncs) {
		fmt.Printf("Max variations of squares is %d\n", len(squareFuncs))
		return
	}
	f = squareFuncs[int(nums)-1]
	s := square.NewSquare(25, f)
	s.Draw()
}
