package main

import (
	"crossword/dictionary"
	"crossword/grid"
	"crossword/utils"
	"fmt"
	"strconv"
	"strings"
)

func (s solver) solve() (map[int]string, map[int]string, grid.Grid) {
	//horizontals, verticals := make(map[int]string), make(map[int]string)
	//filledGrid := grid.Clone()

	// Fill horizontally.
	for i := 0; i < s.g.Height(); i++ {
		_ = s.fillLine(i)
	}

	// Fill holes left
	for j := 0; j < s.g.Width(); j++ {
		var line int
		for _, word := range s.g.WordsInColumn(j) {
			if strings.Contains(word, string(grid.EmptyCell)) {
				regex := strings.ReplaceAll(word, string(grid.EmptyCell), ".")
				match, ok := s.d.ContainsMatch(regex)
				if !ok {
					panic("didn't find a match for final holes: " + regex)
				}
				s.g.FillColumnSegment(line, j, match)
				s.d.Remove(match)
			}
			line += len(word) + 1
		}

	}

	return nil, nil, s.g.Clone()
}

type solver struct {
	g grid.Grid
	d dictionary.Dictionary
}

// TODO handle line not filled
func (s solver) fillLine(line int) bool {
	var inserted []string
	for _, seg := range s.g.FindLineSegments(line) {
		word, ok := s.findCandidate(line, seg.Start, seg.Length)
		if !ok {
			panic("not found line " + strconv.Itoa(line))
		}
		s.g.FillLineSegment(line, seg.Start, word)
		s.d.Remove(word)
		inserted = append(inserted, word)
		fmt.Println("inserting", word, line, seg.Start)
	}

	return true
}

func (s solver) findCandidate(line, column, length int) (string, bool) {
	for word := range s.d {
		regex := s.buildLineSegmentConstraint(line, column, length)
		if !utils.MatchRegex(word, regex) {
			continue
		}

		fmt.Println("verifying candidate", word)
		if s.verifyCandidate(word, line, column) {
			return word, true
		}
	}

	// no candidate found
	return "", false
}

// verifyCandidate checks if a candidate word has matching words on every column.
func (s solver) verifyCandidate(word string, line, column int) bool {
	runed := []rune(word) // take care of multi-character runes
	for j := column; j < s.g.Width() && s.g[line][j] != grid.BlackCell; j++ {
		regex := s.buildColumnConstraint(runed[j-column], line, j)
		fmt.Println("searching on regex", regex)
		if _, ok := s.d.ContainsMatch(regex); !ok {
			fmt.Println("invalid candidate", word, line, column, regex)
			return false
		}
	}
	return true
}

// buildColumnConstraint builds the constraints regex on a column with the candidate letter.
func (s solver) buildColumnConstraint(letter rune, line, column int) string {
	letter = utils.RemoveAccent(letter)
	regex := string(letter)

	// look back first character of word in the column
	for i := line - 1; i >= 0 && s.g[i][column] != grid.BlackCell; i-- {
		// As we fill line by line, previous characters are filled.
		c := utils.RemoveAccent(s.g[i][column])
		if c == grid.EmptyCell {
			c = '.'
		}
		regex = string(c) + regex
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

func (s solver) buildLineSegmentConstraint(line, column, length int) string {
	filled := s.g[line][column : column+length]
	return "^" + strings.ReplaceAll(string(filled), string(grid.EmptyCell), ".") + "$"
}
