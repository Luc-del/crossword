package main

import (
	"crossword/dictionary"
	"crossword/grid"
	v3 "crossword/solver/v3"
	"crossword/utils/logger"
	"fmt"
	"log/slog"
	"strings"
)

func main() {
	logger.Configure(slog.LevelInfo)

	g := grid.New(10, 10)
	initial := g.Clone()
	d := dictionary.NewLengthOrdered("words-shortened.json")

	g.Display()
	h, v, solved := v3.Solve(d, g)

	fmt.Println()
	fmt.Println("----- Initial -----")
	initial.Display()

	fmt.Println()
	fmt.Println("----- Definitions -----")

	fmt.Println("Horizontals:")
	for k, def := range h {
		fmt.Printf("%d: %s\n", k+1, strings.Join(def, " "))
	}

	fmt.Println()
	fmt.Println("Verticals:")
	for k, def := range v {
		fmt.Printf("%s: %s\n", string('A'+rune(k)), strings.Join(def, " "))
	}

	fmt.Println()
	fmt.Println("----- Solution -----")
	solved.Display()
}
