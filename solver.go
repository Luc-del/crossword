package main

import (
	"crossword/dictionary"
	"crossword/grid"
	"fmt"
)

func solve(grid grid.Grid, dict dictionary.Dictionary) (map[int]string, map[int]string, grid.Grid) {
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
	return nil, nil, grid.Clone()
}

type solver struct {
	g grid.Grid
	d dictionary.Dictionary
}

func (s solver) fillLine(word string, line, column int) (string, bool) {
	// find segments FindSegments
	// iter on segments
	// find candidate for each segment
	return word, false
}

func (s solver) findCandidate(line, column int) (string, bool) {
	// iter on dict and verifyCandidate
	// if no candidate return false
	return "", false
}

// verifyCandidate checks if a candidate word can be inserted in the grid.
func (s solver) verifyCandidate(word string, line, column int) bool {
	for j := column; j < s.g.Width() && s.g[line][j] != grid.BlackCell; j++ {
		regex := s.buildConstraint(rune(word[j-line]), line, j)
		if !s.d.ContainsMatch(regex) {
			fmt.Println("invalid candidate", word, line, column, regex)
			return false
		}
	}
	return true
}

// buildConstraint builds the constraints regex on a column with the candidate letter.
func (s solver) buildConstraint(letter rune, line, column int) string {
	regex := string(letter)

	// look back first character of word in the column
	for i := line - 1; i >= 0 && s.g[i][column] != grid.BlackCell; i-- {
		// As we fill line by line, previous characters are filled.
		regex = string(s.g[i][column]) + regex
	}

	// look up last characters of word in the column
	var forwardCount int
	for i := line + 1; i < s.g.Height() && s.g[i][column] != grid.BlackCell; i++ {
		forwardCount++
	}

	if forwardCount > 0 {
		regex += fmt.Sprintf(".{%d}", forwardCount)
	}

	// handle single character: they are not a constraint
	if regex == string(letter) {
		return ""
	}

	return "^" + regex + "$"
}
