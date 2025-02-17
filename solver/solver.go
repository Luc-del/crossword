package solver

import (
	"crossword/dictionary"
	"crossword/grid"
	"log/slog"
	"regexp"
	"strings"
)

type (
	Definitions [][]string

	fill func(*state)
	undo func()
)

type state struct {
	d dictionary.Dictionary
	g grid.Grid

	segments []grid.Segment
	filled   dictionary.Dictionary

	undo undo
}

func Solve(d dictionary.Dictionary, g grid.Grid) (Definitions, Definitions, grid.Grid) {
	root := state{
		d:        d,
		g:        g,
		segments: g.FindAllLineSegments(),
		filled:   make(dictionary.Dictionary),
		undo:     func() { slog.Debug("undo to start") },
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
				if def, ok := root.d[word]; ok {
					root.filled.Add(word, def)
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
		d:        s.d,
		g:        s.g,
		segments: s.segments[segmentIdx+1:],
		filled:   s.filled,
		undo:     func() {},
	}

	lineSegmentFiller(s.segments[segmentIdx].Line, s.segments[segmentIdx].Start, word)(ns)
	for _, f := range fillers {
		f(ns)
	}

	return ns
}

func (s *state) solve() bool {
	if len(s.segments) == 0 {
		return true
	}

	for idx, seg := range s.segments {
		logger := slog.With("line", seg.Line, "column", seg.Start)
		regex := s.buildLineSegmentConstraint(seg)
		if regex == "" { // Empty constraint means the segment is already filled
			logger.Debug("segment skipped")
			continue
		}

		logger = logger.With("regex", regex)
		logger.Debug("looking on segment")

		matcher := regexp.MustCompile(regex)
		for word := range s.d {
			if !matcher.MatchString(word) {
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

		logger.Warn("no candidate, undoing")
		s.undo()
		return false
	}

	panic(",ot here")
}

// verifyCandidate checks if a candidate word has matching words on every column.
func (s *state) verifyCandidate(word string, seg grid.Segment) ([]fill, bool) {
	var fillers []fill
	for j := seg.Start; j < seg.Start+seg.Length; j++ {
		regex := s.buildColumnConstraint(rune(word[j-seg.Start]), seg.Line, j)
		if regex == "" { // Empty constraint means the column is already filled
			continue
		}

		slog.Debug("searching vertically", "regex", regex, "line", seg.Line, "column", j)

		switch match, count := s.d.ContainsMatchN(regex, 2); count {
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

func (s *state) buildLineSegmentConstraint(seg grid.Segment) string {
	filled := s.g[seg.Line][seg.Start : seg.Start+seg.Length]
	regex := "^" + strings.ReplaceAll(string(filled), string(grid.EmptyCell), ".") + "$"
	if !strings.Contains(regex, ".") {
		return ""
	}
	return regex
}

// buildColumnConstraint builds the constraints regex on a column with the candidate letter.
func (s *state) buildColumnConstraint(letter rune, line, column int) string {
	regex := string(letter)

	// look back first character of word in the column
	for i := line - 1; i >= 0 && s.g[i][column] != grid.BlackCell; i-- {
		// As we do line by line, previous characters are filled.
		regex = string(s.g[i][column]) + regex
	}

	// look up last characters of word in the column
	for i := line + 1; i < s.g.Height() && s.g[i][column] != grid.BlackCell; i++ {
		regex += string(s.g[i][column])
	}

	// handle single character: they are not a constraint
	// handle column already filled
	if regex == string(letter) || !strings.Contains(regex, string(grid.EmptyCell)) {
		return ""
	}

	return "^" + strings.ReplaceAll(regex, string(grid.EmptyCell), ".") + "$"
}

func lineSegmentFiller(line, column int, word string) fill {
	return func(newState *state) {
		slog.Info("inserting word horizontally", "word", word, "line", line, "column", column)
		previous := newState.g.FillLineSegment(line, column, word)
		def := newState.d.Pop(word)
		newState.filled.Add(word, def)

		undo := newState.undo
		newState.undo = func() {
			slog.Info("removing word horizontally", "word", word, "line", line, "column", column)
			newState.d.Add(word, def)
			newState.g.FillLineSegment(line, column, previous)
			newState.filled.Remove(word)
			undo()
		}
	}
}

func columnSegmentFiller(line, column int, word string) fill {
	return func(newState *state) {
		slog.Info("inserting word vertically", "word", word, "line", line, "column", column)
		previous := newState.g.FillColumnSegment(line, column, word)
		def := newState.d.Pop(word)
		newState.filled.Add(word, def)

		undo := newState.undo
		newState.undo = func() {
			slog.Info("removing word vertically", "word", word, "line", line, "column", column)
			newState.d.Add(word, def)
			newState.g.FillColumnSegment(line, column, previous)
			newState.filled.Remove(word)
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
