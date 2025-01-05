package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLineSegments(t *testing.T) {
	t.Run("single segment", func(t *testing.T) {
		g := Grid([][]rune{{'_', '_', '_', '_', '_', '_', '_', '_', '_', '_'}})
		assert.Equal(t, []Segment{
			{0, 0, 10},
		},
			g.FindLineSegments(0))
	})

	t.Run("two Xs in a row shouldn't count as a segment", func(t *testing.T) {
		g := Grid([][]rune{{'_', '_', '#', '#', '_', '_', '_', '_', '_', '_'}})
		assert.Equal(t, []Segment{
			{0, 0, 2},
			{0, 4, 6},
		},
			g.FindLineSegments(0))

		g = [][]rune{{'#', '#', '_', '_', '_', '_', '_', '_'}}
		assert.Equal(t, []Segment{
			{0, 2, 6},
		},
			g.FindLineSegments(0))

		g = [][]rune{{'_', '_', '#', '#'}}
		assert.Equal(t, []Segment{
			{0, 0, 2},
		},
			g.FindLineSegments(0))
	})

	t.Run("single letter shouldn't count as a segment", func(t *testing.T) {
		g := Grid([][]rune{{'_', '_', '_', '#', '_', '#', '_', '_', '_', '_'}})
		assert.Equal(t, []Segment{
			{0, 0, 3},
			{0, 6, 4},
		},
			g.FindLineSegments(0))

		g = [][]rune{{'_', '#', '_', '#', '_', '#', '_', '#', '_', '_'}}
		assert.Equal(t, []Segment{
			{0, 8, 2},
		},
			g.FindLineSegments(0))

		g = [][]rune{{'_', '_', '_', '#', '_', '#', '_', '#', '_'}}
		assert.Equal(t, []Segment{
			{0, 0, 3},
		},
			g.FindLineSegments(0))
	})

	t.Run("starting by X", func(t *testing.T) {
		g := Grid([][]rune{{'#', '_', '_', '_', '_', '#', '_', '_', '_', '_'}})
		assert.Equal(t, []Segment{
			{0, 1, 4},
			{0, 6, 4},
		},
			g.FindLineSegments(0))
	})

	t.Run("ending by X", func(t *testing.T) {
		g := Grid([][]rune{{'_', '_', '#', '_', '_', '_', '_', '_', '_', '#'}})
		assert.Equal(t, []Segment{
			{0, 0, 2},
			{0, 3, 6},
		},
			g.FindLineSegments(0))
	})
}

func TestPreviousBlackCellInColumn(t *testing.T) {
	t.Run("no black cell before", func(t *testing.T) {
		g := Grid([][]rune{{'_'}, {'_'}, {BlackCell}})
		idx := g.PreviousBlackCellInColumn(1, 0)
		assert.Equal(t, -1, idx)
	})

	t.Run("black cell first position", func(t *testing.T) {
		g := Grid([][]rune{{BlackCell}, {'_'}, {'_'}})
		idx := g.PreviousBlackCellInColumn(1, 0)
		assert.Equal(t, 0, idx)
	})

	t.Run("black cell right before", func(t *testing.T) {
		g := Grid([][]rune{{'_'}, {BlackCell}, {'_'}})
		idx := g.PreviousBlackCellInColumn(2, 0)
		assert.Equal(t, 1, idx)
	})
}

func TestWordsInColumn(t *testing.T) {
	t.Run("no black cell", func(t *testing.T) {
		g := Grid([][]rune{{'a'}, {'b'}, {'c'}})
		assert.Equal(t, []string{"abc"}, g.WordsInColumn(0))
	})

	t.Run("black cell in middle", func(t *testing.T) {
		g := Grid([][]rune{{'a'}, {BlackCell}, {'c'}})
		assert.Equal(t, []string{"a", "c"}, g.WordsInColumn(0))
	})

	t.Run("black cell first", func(t *testing.T) {
		g := Grid([][]rune{{BlackCell}, {'b'}, {'c'}})
		assert.Equal(t, []string{"bc"}, g.WordsInColumn(0))
	})

	t.Run("black cell last", func(t *testing.T) {
		g := Grid([][]rune{{'a'}, {'b'}, {BlackCell}})
		assert.Equal(t, []string{"ab"}, g.WordsInColumn(0))
	})
}
