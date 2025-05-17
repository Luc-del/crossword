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

	segments  sorted[constrainedSegment]
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
		segments:  newSortedSlice(findAllSegments(g), less),
		usedWords: make(map[string]string),
		undo:      func() { slog.Debug("undone to root") },
	}

	root.solve()

	// TODO definitions

	return Definitions{}, Definitions{}, root.g
}

func findAllSegments(g grid.Grid) []constrainedSegment {
	var (
		res []constrainedSegment
		id  int
	)

	for _, s := range g.FindAllLineSegments() {
		res = append(res, constrainedSegment{
			id:           strconv.Itoa(id),
			isHorizontal: true,
			Position:     s.Position,
			Start:        s.Start,
			Length:       s.Length,
			Constraint:   0,
		})
		id++
	}

	for _, s := range g.FindAllColumnSegments() {
		res = append(res, constrainedSegment{
			id:           strconv.Itoa(id),
			isHorizontal: true,
			Position:     s.Position,
			Start:        s.Start,
			Length:       s.Length,
			Constraint:   0,
		})
		id++
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
		segments:  s.segments[segmentIdx+1:],
		usedWords: s.usedWords,
	}

	if seg.isHorizontal {
		previous := ns.g.FillLineSegment(seg.Position, seg.Start, word)
		def, _ := ns.d.Pop(word)
		ns.usedWords[word] = def
		slog.Debug("inserting word horizontally", "word", word, "line", seg.Position, "column", seg.Start)

		undo := ns.undo
		ns.undo = func() {
			ns.d.Add(word, def)
			ns.g.FillLineSegment(seg.Position, seg.Start, previous)
			delete(ns.usedWords, word)
			slog.Debug("removing word horizontally", "word", word, "line", seg.Position, "column", seg.Start)
			undo()
		}
	} else {
		// TODO fill/unfill colunm
	}

	doCount++
	gridPrinter.Print(slog.Default().With("do", doCount, "undo", undoCount), "new state", ns.g)

	return ns
}
