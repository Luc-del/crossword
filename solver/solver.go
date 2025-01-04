package solver

import (
	"crossword/dictionary"
	"crossword/grid"
	"log/slog"
	"regexp"
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

func (s *Solver) Solve() (Definitions, Definitions, grid.Grid) {
	// Fill horizontally.
	for i := 0; i < s.g.Height(); i++ {
		_ = s.fillLine(i)
	}

	// Fill holes left
	for j := 0; j < s.g.Width(); j++ {
		var line int
		for _, word := range s.g.WordsInColumn(j) {
			switch {
			case len(word) == 1:
			// pass
			case strings.Contains(word, string(grid.EmptyCell)):
				regex := strings.ReplaceAll(word, string(grid.EmptyCell), ".")
				match, ok := s.d.ContainsMatch(regex)
				if !ok {
					panic("didn't find a match for final holes: " + regex)
				}
				s.fillColumnSegment(line, j, match)
			default:
				if def, ok := s.d[word]; ok {
					s.filled[word] = def
					s.d.Remove(word)
				}
			}
			line += len(word) + 1
		}
	}

	h, v := s.extractDefinitions()
	return h, v, s.g
}

type fillAction func()

// TODO handle line not filled
func (s *Solver) fillLine(line int) bool {
	for _, seg := range s.g.FindLineSegments(line) {
		fillers, ok := s.findCandidate(line, seg.Start, seg.Length)
		if !ok {
			panic("not found line/column " + strconv.Itoa(line) + " " + strconv.Itoa(seg.Start))
		}

		for _, f := range fillers {
			f()
		}
	}

	return true
}

func (s Solver) findCandidate(line, column, length int) ([]fillAction, bool) {
	regex := s.buildLineSegmentConstraint(line, column, length)
	if regex == "" {
		// Empty constraint means the line is already filled
		return nil, true
	}

	matcher := regexp.MustCompile(regex)
	for word := range s.d {
		if !matcher.MatchString(word) {
			continue
		}

		slog.Debug("verifying candidate", "word", word)
		fillers, ok := s.verifyCandidate(word, line, column)
		if ok {
			return append(fillers, func() {
				s.fillLineSegment(line, column, word)
			}), true
		}
	}

	// no candidate found
	return nil, false
}

// verifyCandidate checks if a candidate word has matching words on every column.
func (s *Solver) verifyCandidate(word string, line, column int) ([]fillAction, bool) {
	var fillers []fillAction
	for j := column; j < s.g.Width() && s.g[line][j] != grid.BlackCell; j++ {
		regex := s.buildColumnConstraint(rune(word[j-column]), line, j)
		if regex == "" { // no constraint
			continue
		}

		logger := slog.With("regex", regex, "line", line, "column", j)
		logger.Debug("searching vertically")

		switch match, count := s.d.ContainsMatchN(regex, 2); count {
		case 0:
			logger.Debug("invalid candidate", "word", word)
			return nil, false
		case 1:
			fillers = append(fillers, func() {
				start := s.g.PreviousBlackCellInColumn(line, j)
				s.fillColumnSegment(start+1, j, match)
			})
		}
	}
	return fillers, true
}

// buildColumnConstraint builds the constraints regex on a column with the candidate letter.
func (s *Solver) buildColumnConstraint(letter rune, line, column int) string {
	regex := string(letter)

	// look back first character of word in the column
	for i := line - 1; i >= 0 && s.g[i][column] != grid.BlackCell; i-- {
		// As we fill line by line, previous characters are filled.
		regex = string(s.g[i][column]) + regex
	}

	// look up last characters of word in the column
	for i := line + 1; i < s.g.Height() && s.g[i][column] != grid.BlackCell; i++ {
		regex += string(s.g[i][column])
	}

	switch {
	// handle single character: they are not a constraint
	case regex == string(letter),
		// handle column already filled
		!strings.Contains(regex, string(grid.EmptyCell)):
		return ""
	}

	return "^" + strings.ReplaceAll(regex, string(grid.EmptyCell), ".") + "$"
}

func (s *Solver) buildLineSegmentConstraint(line, column, length int) string {
	filled := s.g[line][column : column+length]
	regex := "^" + strings.ReplaceAll(string(filled), string(grid.EmptyCell), ".") + "$"
	if !strings.Contains(regex, ".") {
		return ""
	}
	return regex
}

type Definitions [][]string

func (s *Solver) extractDefinitions() (Definitions, Definitions) {
	h, v := make(Definitions, s.g.Height()), make(Definitions, s.g.Width())

	for i := range s.g {
		var def []string
		for _, w := range strings.Split(string(s.g[i]), string(grid.BlackCell)) {
			if len(w) <= 1 {
				continue
			}

			d, ok := s.filled[w]
			if !ok {
				// Extract definitions for autofilled lines
				d = s.d[w]
				s.d.Remove(w)
			}
			def = append(def, d)
		}
		h[i] = def
	}

	for j := range s.g.Width() {
		var words []rune
		for i := range s.g.Height() {
			words = append(words, s.g[i][j])
		}

		for _, w := range strings.Split(string(words), string(grid.BlackCell)) {
			if len(w) <= 1 {
				continue
			}
			v[j] = append(v[j], s.filled[w])
		}
	}

	return h, v
}

func (s *Solver) fillColumnSegment(line, column int, word string) {
	slog.Info("inserting word vertically", "word", word, "line", line, "column", column)
	s.g.FillColumnSegment(line, column, word)
	s.filled[word] = s.d[word]
	s.d.Remove(word)
}

func (s *Solver) fillLineSegment(line, column int, word string) {
	slog.Info("inserting word horizontally", "word", word, "line", line, "column", column)
	s.g.FillLineSegment(line, column, word)
	s.filled[word] = s.d[word]
	s.d.Remove(word)
}
