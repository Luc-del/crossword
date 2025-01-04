package main

import (
	"crossword/dictionary"
	"crossword/grid"
	"crossword/solver"
	"crossword/utils/logger"
	"fmt"
	"log/slog"
	"strings"
)

func main() {
	logger.Configure(slog.LevelError)

	g := grid.ExampleGrid.Clone()
	g.FillLineSegment(4, 4, "NEO")
	g.FillLineSegment(4, 8, "OR")
	d := dictionary.NewExample()

	s := solver.New(d, g)
	g.Display()

	fmt.Println()
	fmt.Println("----- Definitions -----")

	h, v, solved := s.Solve()
	fmt.Println("Horizontals:")
	for k, def := range h {
		fmt.Printf("%d: %s\n", k+1, strings.Join(def, " "))
	}

	fmt.Println("Verticals:")
	for k, def := range v {
		fmt.Printf("%s: %s\n", string('A'+rune(k)), strings.Join(def, " "))
	}

	fmt.Println()
	fmt.Println("----- Solution -----")
	solved.Display()
}
