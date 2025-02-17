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
	d := dictionary.NewExample()

	g.Display()
	h, v, solved := solver.Solve(d, g)

	fmt.Println()
	fmt.Println("----- Definitions -----")

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
