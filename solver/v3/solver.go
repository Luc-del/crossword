// Package v3 solves a grid by iterating on line segments searching by pattern.
package v3

import (
	"crossword/grid"
	"crossword/matcher"
	"crossword/printer"
	"log/slog"
	"strings"
	"time"
)

var gridPrinter = printer.New(100)

type dictionary interface {
	Remove(string)
	Pop(string) (string, bool)
	Add(word, def string)
	ContainsMatch(regex string) (string, bool)
	ContainsPatternN(pattern []rune, atLeast int) (string, int)
	Registry(length int) map[string]string
}

type (
	Definitions [][]string

	fill func(*state)
	undo func()
)

type state struct {
	depth int

	d dictionary
	g grid.Grid

	segments  []grid.Segment
	usedWords map[string]string

	undo undo
}

var (
	doCount   int
	undoCount int
)

func Solve(d dictionary, g grid.Grid) (Definitions, Definitions, grid.Grid) {
	start := time.Now()
	defer func() { slog.Info("time monitoring", "elapsed", time.Since(start).String()) }()

	root := state{
		depth:     0,
		d:         d,
		g:         g,
		segments:  g.FindAllLineSegments(),
		usedWords: make(map[string]string),
		undo:      func() { slog.Debug("undone to root") },
	}

	root.solve()

	// Fill holes left
	for j := 0; j < g.Width(); j++ {
		var line int
		for _, word := range g.WordsInColumn(j) {
			switch {
			case len(word) == 1:
			// pass
			case strings.Contains(word, string(grid.EmptyCell)):
				regex := strings.ReplaceAll(word, string(grid.EmptyCell), ".")
				match, ok := d.ContainsMatch(regex)
				if !ok {
					panic("didn't find a match for final holes: " + regex)
				}
				columnSegmentFiller(line, j, match)(&root)
			default:
				if def, ok := root.d.Pop(word); ok {
					root.usedWords[word] = def
					root.d.Remove(word)
				}
			}
			line += len(word) + 1
		}
	}

	h, v := root.extractDefinitions()

	return h, v, root.g
}

func (s *state) mutate(segmentIdx int, word string, fillers []fill) *state {
	ns := &state{
		depth:     s.depth + 1,
		d:         s.d,
		g:         s.g,
		segments:  s.segments[segmentIdx+1:],
		usedWords: s.usedWords,
		undo:      func() {},
	}

	lineSegmentFiller(s.segments[segmentIdx].Line, s.segments[segmentIdx].Start, word)(ns)
	for _, f := range fillers {
		f(ns)
	}

	doCount++
	gridPrinter.Print(slog.Default().With("do", doCount, "undo", undoCount), "new state", ns.g)

	return ns
}

func (s *state) solve() bool {
	if len(s.segments) == 0 {
		return true
	}

	for idx, seg := range s.segments {
		logger := slog.With("line", seg.Line, "column", seg.Start)
		pattern := s.buildLineSegmentConstraint(seg)
		if !strings.Contains(string(pattern), string(grid.EmptyCell)) { // Line is already filled
			logger.Debug("segment skipped")
			continue
		}

		logger = logger.With("pattern", string(pattern))
		logger.Debug("looking on segment")

		for word := range s.d.Registry(len(pattern)) {
			if !matcher.Pattern(word, pattern) {
				continue
			}

			logger := logger.With("word", word)
			logger.Debug("verifying word")
			fillers, ok := s.verifyCandidate(word, seg)
			if !ok {
				logger.Debug("invalid candidate")
				continue
			}

			newState := s.mutate(idx, word, fillers)
			if newState.solve() {
				return true
			}
		}

		logger.Debug("no candidate, undoing")
		s.undo()
		undoCount++
		gridPrinter.Print(slog.Default().With("do", doCount, "undo", undoCount), "new state after undo", s.g)

		return false
	}

	panic("not here")
}

// verifyCandidate checks if a candidate word has matching words on every column.
func (s *state) verifyCandidate(word string, seg grid.Segment) ([]fill, bool) {
	var fillers []fill
	for j := seg.Start; j < seg.Start+seg.Length; j++ {
		pattern := s.findColumnConstraint(rune(word[j-seg.Start]), seg.Line, j)
		// handle single character: they are not a constraint
		// handle column already filled
		if len(pattern) == 1 || !strings.Contains(string(pattern), string(grid.EmptyCell)) {
			continue
		}

		slog.Debug("searching vertically", "pattern", string(pattern), "line", seg.Line, "column", j)

		switch match, count := s.d.ContainsPatternN(pattern, 2); count { // TODO exclude current word
		case 0:
			return nil, false
		case 1:
			fillers = append(fillers, func(newState *state) {
				start := s.g.PreviousBlackCellInColumn(seg.Line, j)
				columnSegmentFiller(start+1, j, match)(newState)
			})
		}
	}
	return fillers, true
}

func (s *state) buildLineSegmentConstraint(seg grid.Segment) []rune {
	return s.g[seg.Line][seg.Start : seg.Start+seg.Length]
}

func (s *state) findColumnConstraint(letter rune, line, column int) []rune {
	pattern := []rune{letter}

	// look back first character of word in the column
	for i := line - 1; i >= 0 && s.g[i][column] != grid.BlackCell; i-- {
		// As we do line by line, previous characters are usedWords.
		pattern = append([]rune{s.g[i][column]}, pattern...)
	}

	// look up last characters of word in the column
	for i := line + 1; i < s.g.Height() && s.g[i][column] != grid.BlackCell; i++ {
		pattern = append(pattern, s.g[i][column])
	}

	return pattern
}

func lineSegmentFiller(line, column int, word string) fill {
	return func(newState *state) {
		previous := newState.g.FillLineSegment(line, column, word)
		def, _ := newState.d.Pop(word)
		newState.usedWords[word] = def
		slog.Debug("inserting word horizontally", "word", word, "line", line, "column", column)

		undo := newState.undo
		newState.undo = func() {
			newState.d.Add(word, def)
			newState.g.FillLineSegment(line, column, previous)
			delete(newState.usedWords, word)
			slog.Debug("removing word horizontally", "word", word, "line", line, "column", column)
			undo()
		}
	}
}

func columnSegmentFiller(line, column int, word string) fill {
	return func(newState *state) {
		previous := newState.g.FillColumnSegment(line, column, word)
		def, _ := newState.d.Pop(word)
		newState.usedWords[word] = def
		slog.Debug("inserting word vertically", "word", word, "line", line, "column", column)

		undo := newState.undo
		newState.undo = func() {
			newState.d.Add(word, def)
			newState.g.FillColumnSegment(line, column, previous)
			delete(newState.usedWords, word)
			slog.Debug("removing word vertically", "word", word, "line", line, "column", column)
			undo()
		}
	}
}

func (s *state) extractDefinitions() (Definitions, Definitions) {
	h, v := make(Definitions, s.g.Height()), make(Definitions, s.g.Width())

	for i := range s.g {
		var def []string
		for _, w := range strings.Split(string(s.g[i]), string(grid.BlackCell)) {
			if len(w) <= 1 {
				continue
			}

			d, ok := s.usedWords[w]
			if !ok {
				// Extract definitions for autofilled lines
				d, _ = s.d.Pop(w)
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
			v[j] = append(v[j], s.usedWords[w])
		}
	}

	return h, v
}
