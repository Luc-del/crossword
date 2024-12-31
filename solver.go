package main

import "crossword/grid"

//func solve(grid grid.Grid, words []Word) (map[int]string, map[int]string, [][]rune) {
//	//horizontals := make(map[int]string)
//	//verticals := make(map[int]string)
//	//
//	//filledGrid := grid.Clone()
//	//
//	//// Fill horizontally.
//	//for i := 0; i < gridSize; i++ {
//	//	segments := grid.FindSegments(filledGrid[i])
//	//	for _, seg := range segments {
//	//
//	//	}
//	//}
//
//}

//func wordIsOKOnLineSegment(g grid.Grid, words []Word, word Word, line, column int) bool {
//	for j := column; g[line][j] != grid.BlackCell; j++ {
//
//	}
//}

func checkConstraints

// columnRegex builds the constraints regex on a column with the candidate letter.
func buildConstraint(g grid.Grid, letter rune, line, column int) string {
	regex := string(letter)

	// look back first character of word in the column
	for i := line - 1; i >= 0 && g[i][column] != grid.BlackCell; i-- {
		// As we fill line by line, previous characters are filled.
		regex = string(g[i][column]) + regex
	}

	// look up last characters of word in the column
	for i := line + 1; i < len(g) && g[i][column] != grid.BlackCell; i++ {
		regex += "*"
	}

	return regex
}
