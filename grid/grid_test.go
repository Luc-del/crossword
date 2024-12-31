package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid_FindLineSegments(t *testing.T) {
	t.Run("single segment", func(t *testing.T) {
		g := Grid([][]rune{{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'}})
		assert.Equal(t, []Segment{
			{0, 10},
		},
			g.FindLineSegments(0))
	})

	t.Run("two Xs in a row shouldn't count as a segment", func(t *testing.T) {
		g := Grid([][]rune{{'.', '.', 'X', 'X', '.', '.', '.', '.', '.', '.'}})
		assert.Equal(t, []Segment{
			{0, 2},
			{4, 6},
		},
			g.FindLineSegments(0))

		g = [][]rune{{'X', 'X', '.', '.', '.', '.', '.', '.'}}
		assert.Equal(t, []Segment{
			{2, 6},
		},
			g.FindLineSegments(0))

		g = [][]rune{{'.', '.', 'X', 'X'}}
		assert.Equal(t, []Segment{
			{0, 2},
		},
			g.FindLineSegments(0))
	})

	t.Run("single letter shouldn't count as a segment", func(t *testing.T) {
		g := Grid([][]rune{{'.', '.', '.', 'X', '.', 'X', '.', '.', '.', '.'}})
		assert.Equal(t, []Segment{
			{0, 3},
			{6, 4},
		},
			g.FindLineSegments(0))

		g = [][]rune{{'.', 'X', '.', 'X', '.', 'X', '.', 'X', '.', '.'}}
		assert.Equal(t, []Segment{
			{8, 2},
		},
			g.FindLineSegments(0))

		g = [][]rune{{'.', '.', '.', 'X', '.', 'X', '.', 'X', '.'}}
		assert.Equal(t, []Segment{
			{0, 3},
		},
			g.FindLineSegments(0))
	})

	t.Run("starting by X", func(t *testing.T) {
		g := Grid([][]rune{{'X', '.', '.', '.', '.', 'X', '.', '.', '.', '.'}})
		assert.Equal(t, []Segment{
			{1, 4},
			{6, 4},
		},
			g.FindLineSegments(0))
	})

	t.Run("ending by X", func(t *testing.T) {
		g := Grid([][]rune{{'.', '.', 'X', '.', '.', '.', '.', '.', '.', 'X'}})
		assert.Equal(t, []Segment{
			{0, 2},
			{3, 6},
		},
			g.FindLineSegments(0))
	})
}
