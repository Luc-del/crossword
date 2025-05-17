// Package v4 solves a grid by iterating on most constrained line segments searching by pattern.
package v4

import (
	"crossword/grid"
	"crossword/matcher"
	"crossword/printer"
	"log/slog"
	"strconv"
	"time"
)

var (
	gridPrinter = printer.New(500)

	doCount   int
	undoCount int
)

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

	undo func()
)

type state struct {
	depth int

	d dictionary
	g grid.Grid

	segments  segmentRepository
	usedWords map[string]string

	undo undo
}

func Solve(d dictionary, g grid.Grid) (Definitions, Definitions, grid.Grid) {
	start := time.Now()
	defer func() { slog.Info("time monitoring", "elapsed", time.Since(start).String()) }()

	root := state{
		depth:     0,
		d:         d,
		g:         g,
		segments:  newSegmentRepository(findAllSegments(g)),
		usedWords: make(map[string]string),
		undo:      func() { slog.Debug("undone to root") },
	}

	root.solve()

	// TODO definitions

	return Definitions{}, Definitions{}, root.g
}

func findAllSegments(g grid.Grid) map[string]constrainedSegment {
	var (
		res     = make(map[string]constrainedSegment)
		counter int
	)

	for _, s := range g.FindAllLineSegments() {
		id := strconv.Itoa(counter)
		res[id] = constrainedSegment{
			id:           id,
			isHorizontal: true,
			Position:     s.Position,
			Start:        s.Start,
			Length:       s.Length,
			Constraint:   0,
		}
		counter++
	}

	for _, s := range g.FindAllColumnSegments() {
		id := strconv.Itoa(counter)
		res[id] = constrainedSegment{
			id:           id,
			isHorizontal: false,
			Position:     s.Position,
			Start:        s.Start,
			Length:       s.Length,
			Constraint:   0,
		}
		counter++
	}

	return res
}

func (s *state) solve() bool {
	seg := s.segments.GetMax()
	if seg.Constraint == 0 {
		return true
	}

	pattern := s.buildSegmentConstraint(seg)

	logger := slog.With("direction", seg.isHorizontal, "start", seg.Start, "length", seg.Length, "pattern", string(pattern))
	logger.Debug("looking on segment")

	for word := range s.d.Registry(len(pattern)) {
		if !matcher.Pattern(word, pattern) {
			continue
		}

		newState := s.mutate(seg, word)
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

func (s *state) buildSegmentConstraint(seg constrainedSegment) []rune {
	if seg.isHorizontal {
		return s.g[seg.Position][seg.Start : seg.Start+seg.Length]
	}

	res := make([]rune, seg.Length)
	for i := 0; i < seg.Length; i++ {
		res[i] = s.g[seg.Start+i][seg.Position]
	}
	return res
}

func (s *state) mutate(seg constrainedSegment, word string) *state {
	ns := &state{
		depth:     s.depth + 1,
		d:         s.d,
		g:         s.g,
		segments:  s.segments,
		usedWords: s.usedWords,
	}

	if seg.isHorizontal {
		previous := ns.g.FillLineSegment(seg.Position, seg.Start, word)
		def, _ := ns.d.Pop(word)
		ns.usedWords[word] = def
		for j := seg.Start; j < seg.Start+seg.Length; j++ {
			ns.segments.IncrementConstraint(seg.Position, j, true)
		}
		slog.Debug("inserting word horizontally", "word", word, "line", seg.Position, "column", seg.Start)

		undo := ns.undo
		ns.undo = func() {
			ns.d.Add(word, def)
			ns.g.FillLineSegment(seg.Position, seg.Start, previous)
			delete(ns.usedWords, word)
			for j := seg.Start; j < seg.Start+seg.Length; j++ {
				ns.segments.DecrementConstraint(seg.Position, j, true)
			}
			slog.Debug("removing word horizontally", "word", word, "line", seg.Position, "column", seg.Start)
			undo()
		}
	} else {
		previous := ns.g.FillColumnSegment(seg.Start, seg.Position, word)
		def, _ := ns.d.Pop(word)
		ns.usedWords[word] = def
		for i := seg.Start; i < seg.Start+seg.Length; i++ {
			ns.segments.IncrementConstraint(i, seg.Position, false)
		}
		slog.Debug("inserting word vertically", "word", word, "line", seg.Start, "column", seg.Position)

		undo := ns.undo
		ns.undo = func() {
			ns.d.Add(word, def)
			ns.g.FillColumnSegment(seg.Start, seg.Position, previous)
			delete(ns.usedWords, word)
			for i := seg.Start; i < seg.Start+seg.Length; i++ {
				ns.segments.DecrementConstraint(i, seg.Position, false)
			}
			slog.Debug("removing word vertically", "word", word, "line", seg.Start, "column", seg.Position)
			undo()
		}
	}

	doCount++
	gridPrinter.Print(slog.Default().With("do", doCount, "undo", undoCount), "new state", ns.g)

	return ns
}
