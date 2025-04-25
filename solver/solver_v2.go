package solver

import (
	"crossword/grid"
	"fmt"
	"log/slog"
	"regexp"
	"strings"
	"time"
)

type cursor struct {
	line, column int
}

type statev2 struct {
	depth int
	d     dictionary
	g     grid.Grid

	c         cursor
	usedWords map[string]string

	undo undo
}

// SolveFromEmptyGrid finds words and black cells to fit a given grid.
func SolveFromEmptyGrid(d dictionary, g grid.Grid) (Definitions, Definitions, grid.Grid) {
	start := time.Now()
	defer func() { slog.Info("time monitoring", "elapsed", time.Since(start).String()) }()

	root := statev2{
		depth:     0,
		d:         d,
		g:         g,
		c:         cursor{line: 0, column: 0},
		usedWords: make(map[string]string),
		undo:      func() { slog.Debug("undone to root") },
	}

	root.solve()

	return Definitions{}, Definitions{}, root.g
}

func (s *statev2) solve() bool {
	logger := slog.With("line", s.c.line, "column", s.c.column, "depth", s.depth, "completion", s.g.CompletionState())

	switch s.c.column {
	case s.g.Width(), s.g.Width() - 1: // Line is usedWords or single character left
		if s.c.line == s.g.Height()-1 {
			logger.Debug("end of grid reached")
			return true
		}

		logger.Debug("going to new line")
		newState := statev2{
			depth:     s.depth,
			d:         s.d,
			g:         s.g,
			c:         cursor{line: s.c.line + 1, column: 0},
			usedWords: s.usedWords,
			undo:      func() { slog.Debug("undone to root") },
		}
		return newState.solve()
	}

	regex := s.buildLineSegmentConstraint()
	logger = logger.With("regex", regex)
	logger.Debug("looking on segment")

	matcher := regexp.MustCompile(regex)
	for word := range s.d.Registry(regex) {
		if !matcher.MatchString(word) {
			continue
		}

		logger := logger.With("word", word)
		logger.Debug("verifying word")
		if !s.verifyCandidate(word) {
			logger.Debug("invalid candidate")
			continue
		}

		newState := s.mutate(word)
		if newState.solve() {
			return true
		}
	}

	logger.Warn("no candidate, undoing")
	s.undo()
	return false

}

// verifyCandidate checks if a candidate word has matching words on every column.
func (s *statev2) verifyCandidate(word string) bool {
	for j := s.c.column; j < s.c.column+len(word); j++ {
		regex := s.buildColumnConstraint(rune(word[j-s.c.column]), s.c.line, j)
		if regex == "" { // Empty constraint means the column is already usedWords
			continue
		}

		slog.Debug("searching vertically", "regex", regex, "line", s.c.line, "column", j)
		if _, count := s.d.ContainsMatchN(regex, 1); count == 0 { // TODO exclude current word
			return false
		}
	}

	return true
}

func (s *statev2) buildLineSegmentConstraint() string {
	return fmt.Sprintf("^.{2,%d}$", s.g.Width()-s.c.column)
}

// buildColumnConstraint builds the constraints regex on a column with the candidate letter.
func (s *statev2) buildColumnConstraint(letter rune, line, column int) string {
	regex := string(letter)

	// look back first character of word in the column
	for i := line - 1; i >= 0 && s.g[i][column] != grid.BlackCell; i-- {
		// As we do line by line, previous characters are usedWords.
		regex = string(s.g[i][column]) + regex
	}

	regex = fmt.Sprintf("%s.{%d,%d}", regex, 0, s.g.Height()-line-1)

	// handle single character: they are not a constraint
	if regex == string(letter) {
		return ""
	}

	return "^" + strings.ReplaceAll(regex, string(grid.EmptyCell), ".") + "$"
}

func (s *statev2) mutate(word string) *statev2 {
	ns := &statev2{
		depth:     s.depth + 1,
		d:         s.d,
		g:         s.g,
		c:         cursor{line: s.c.line, column: s.c.column + len(word)}, // TODO probability of black cell on the left of the grid
		usedWords: s.usedWords,
		undo:      func() {},
	}

	ns.g.FillLineSegment(s.c.line, s.c.column, word)
	if ns.c.column < s.g.Width() {
		ns.g[ns.c.line][ns.c.column] = grid.BlackCell
		ns.c.column++
	}
	def, _ := ns.d.Pop(word)
	ns.usedWords[word] = def
	slog.Info("inserting word horizontally", "word", word, "line", s.c.line, "column", s.c.column, "completion", ns.g.CompletionState())
	ns.g.Print()
	time.Sleep(1 * time.Second)

	ns.undo = func() {
		ns.d.Add(word, def)
		ns.g.EmptyLineSegment(s.c.line, s.c.column, len(word))
		delete(ns.usedWords, word)
		slog.Info("removing word horizontally", "word", word, "line", s.c.line, "column", s.c.column, "completion", ns.g.CompletionState())
		ns.g.Print()
		time.Sleep(1 * time.Second)
	}

	return ns
}
