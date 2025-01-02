package solver

import (
	"crossword/dictionary"
	"crossword/grid"
	"crossword/utils"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct {
	d, filled dictionary.Dictionary
	g         grid.Grid
}

func New(d dictionary.Dictionary, g grid.Grid) Solver {
	return Solver{
		d:      d,
		filled: make(dictionary.Dictionary),
		g:      g,
	}
}

func (s Solver) Solve() (Definitions, Definitions, grid.Grid) {
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
				s.filled[match] = s.d[match]
				s.d.Remove(match)
			}
			line += len(word) + 1
		}
	}

	h, v := s.extractDefinitions()
	return h, v, s.g
}

// TODO handle line not filled
func (s Solver) fillLine(line int) bool {
	var inserted []string
	for _, seg := range s.g.FindLineSegments(line) {
		match, ok := s.findCandidate(line, seg.Start, seg.Length)
		if !ok {
			panic("not found line " + strconv.Itoa(line))
		}
		s.g.FillLineSegment(line, seg.Start, match)
		s.filled[match] = s.d[match]
		s.d.Remove(match)
		inserted = append(inserted, match)
		fmt.Println("inserting", match, line, seg.Start)
	}

	return true
}

func (s Solver) findCandidate(line, column, length int) (string, bool) {
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
func (s Solver) verifyCandidate(word string, line, column int) bool {
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
func (s Solver) buildColumnConstraint(letter rune, line, column int) string {
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

func (s Solver) buildLineSegmentConstraint(line, column, length int) string {
	filled := s.g[line][column : column+length]
	return "^" + strings.ReplaceAll(string(filled), string(grid.EmptyCell), ".") + "$"
}

type Definitions [][]string

func (s Solver) extractDefinitions() (Definitions, Definitions) {
	h, v := make(Definitions, s.g.Height()), make(Definitions, s.g.Width())

	for i := range s.g {
		words := strings.Split(string(s.g[i]), string(grid.BlackCell))
		for k, w := range words {
			words[k] = s.filled[w]
		}
		h[i] = words
	}

	for j := range s.g[0] {
		words := strings.Split(string(s.g[j]), string(grid.BlackCell))
		for k, w := range words {
			words[k] = s.filled[w]
		}
		v[j] = words
	}

	return h, v
}
