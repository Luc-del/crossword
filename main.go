package main

import (
	"crossword/dictionary"
	"crossword/grid"
	"crossword/solver"
	"fmt"
	"strings"
)

func main() {
	g := grid.ExampleGrid.Clone()
	g.FillLineSegment(4, 4, "neo")
	g.FillLineSegment(4, 8, "or")
	d := dictionary.NewExample()

	s := solver.New(d, g)
	g.Print()

	h, v, solved := s.Solve()
	fmt.Println("Horizontals:")
	for k, def := range h {
		fmt.Printf("%d: %s\n", k+1, strings.Join(def, " "))
	}

	fmt.Println("Verticals:")
	for k, def := range v {
		fmt.Printf("%s: %s\n", string('A'+rune(k)), strings.Join(def, " "))
	}

	solved.Uppercase().Print()
}
