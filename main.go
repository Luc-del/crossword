package main

import (
	"fmt"
)

const gridSize = 10

func main() {
	// Exemple de grille fournie
	grid := [][]rune{
		{'.', '.', '.', 'X', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'X', '.', '.', '.', '.', '.', '.'},
		{'X', '.', '.', '.', '.', 'X', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', 'X', '.', '.', '.', '.'},
		{'.', '.', 'X', 'X', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', 'X', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', 'X', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', 'X'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'X', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	horizontals, verticals, filledGrid := solveCrossword(grid, dictionary)

	fmt.Println("Grille initiale :")
	printGrid(grid)

	fmt.Println("\nSolution complète :")
	printGrid(filledGrid)

	fmt.Println("\nDéfinitions horizontales :")
	for k, def := range horizontals {
		fmt.Printf("Ligne %d: %s\n", k, def)
	}

	fmt.Println("\nDéfinitions verticales :")
	for k, def := range verticals {
		fmt.Printf("Colonne %d: %s\n", k, def)
	}
}
