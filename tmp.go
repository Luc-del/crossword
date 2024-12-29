package main

import "fmt"

// Résolution de la grille
func solveCrossword(grid [][]rune, dictionary []Word) (map[int]string, map[int]string, [][]rune) {
	horizontals := make(map[int]string)
	verticals := make(map[int]string)

	usedWords := make(map[string]bool)
	filledGrid := cloneGrid(grid)

	// Remplissage horizontal
	for i := 0; i < gridSize; i++ {
		segments := findSegments(filledGrid[i])
		fmt.Println("CORNICHON horizontal", i, segments)
		for _, seg := range segments {
			word := getFittingWord(seg.length, dictionary, usedWords)
			if word != nil {
				placeWord(filledGrid, word.Text, i, seg.start, true)
				horizontals[i] = word.Definition
				usedWords[word.Text] = true
			}
		}
	}

	// Remplissage vertical
	for j := 0; j < gridSize; j++ {
		column := extractColumn(filledGrid, j)
		segments := findSegments(column)
		for _, seg := range segments {
			word := getFittingWord(seg.length, dictionary, usedWords)
			if word != nil {
				placeWord(filledGrid, word.Text, seg.start, j, false)
				verticals[j] = word.Definition
				usedWords[word.Text] = true
			}
		}
	}

	return horizontals, verticals, filledGrid
}

func cloneGrid(grid [][]rune) [][]rune {
	clone := make([][]rune, len(grid))
	for i := range grid {
		clone[i] = make([]rune, len(grid[i]))
		copy(clone[i], grid[i])
	}
	return clone
}

// Trouve les segments de longueur disponibles dans une ligne ou colonne
func findSegments(line []rune) []struct {
	start, length int
} {
	var segments []struct {
		start, length int
	}
	start := -1
	for i, char := range line {
		if char == '.' {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				segments = append(segments, struct{ start, length int }{start, i - start})
				start = -1
			}
		}
	}
	if start != -1 {
		segments = append(segments, struct{ start, length int }{start, len(line) - start})
	}
	return segments
}

// Sélectionne un mot adapté à la longueur demandée
func getFittingWord(length int, dictionary []Word, usedWords map[string]bool) *Word {
	for _, word := range dictionary {
		if len(word.Text) == length && !usedWords[word.Text] {
			return &word
		}
	}
	return nil
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

// Affiche la grille
func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}
}
