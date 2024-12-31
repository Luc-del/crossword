package main

import (
	"crossword/dictionary"
	"crossword/grid"
)

// Résolution de la grille
func solveCrossword(grid grid.Grid, words []dictionary.Word) (map[int]string, map[int]string, grid.Grid) {
	horizontals := make(map[int]string)
	verticals := make(map[int]string)

	//usedWords := make(map[string]bool)
	filledGrid := grid.Clone()

	//// Remplissage horizontal
	//for i := 0; i < gridSize; i++ {
	//	segments := findSegments(filledGrid[i])
	//	for _, seg := range segments {
	//		if hasFittingWord(seg.length, words) {
	//			placeWord(filledGrid, word.Text, i, seg.start, true)
	//			horizontals[i] = word.Definition
	//			usedWords[word.Text] = true
	//		}
	//	}
	//}
	//
	//// Remplissage vertical
	//for j := 0; j < gridSize; j++ {
	//	column := extractColumn(filledGrid, j)
	//	segments := findSegments(column)
	//	for _, seg := range segments {
	//		word := getFittingWord(seg.length, words, usedWords)
	//		if word != nil {
	//			placeWord(filledGrid, word.Text, seg.start, j, false)
	//			verticals[j] = word.Definition
	//			usedWords[word.Text] = true
	//		}
	//	}
	//}

	return horizontals, verticals, filledGrid
}

// Place un mot dans la grille
func placeWord(grid [][]rune, word string, x, y int, horizontal bool) {
	for i, char := range word {
		if horizontal {
			grid[x][y+i] = char
		} else {
			grid[x+i][y] = char
		}
	}
}

// Extrait une colonne de la grille
func extractColumn(grid [][]rune, col int) []rune {
	column := make([]rune, gridSize)
	for i := 0; i < gridSize; i++ {
		column[i] = grid[i][col]
	}
	return column
}
