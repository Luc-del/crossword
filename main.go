package main

import (
	"crossword/dictionary"
	"crossword/grid"
	"fmt"
)

const gridSize = 10

func main() {

	g := grid.NewGrid()

	horizontals, verticals, filledGrid := solve(g, dictionary.NewDefault())

	fmt.Println("Grille initiale :")
	g.Print()

	fmt.Println("\nSolution complète :")
	filledGrid.Print()

	fmt.Println("\nDéfinitions horizontales :")
	for k, def := range horizontals {
		fmt.Printf("Ligne %d: %s\n", k, def)
	}

	fmt.Println("\nDéfinitions verticales :")
	for k, def := range verticals {
		fmt.Printf("Colonne %d: %s\n", k, def)
	}
}
