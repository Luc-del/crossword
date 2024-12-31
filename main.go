package main

import (
	grid2 "crossword/grid"
	"fmt"
)

const gridSize = 10

func main() {

	grid := grid2.NewGrid()

	horizontals, verticals, filledGrid := solveCrossword(grid, words)

	fmt.Println("Grille initiale :")
	grid.Print()

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
