package v3

import (
	"crossword/example"
	"crossword/grid"
	"crossword/utils/logger"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	logger.Configure(slog.LevelDebug)

	d := example.Dictionary()
	d.Add("tes", "several possibilities until the end, make a choice")

	h, v, res := Solve(d, example.Grid)

	// Two possibilities left
	if res[2][9] == 'S' {
		example.Solved[2][9] = 'S'
		example.Vertical[9][0] = "several possibilities until the end, make a choice"
	}

	assert.Equal(t, grid.Grid(example.Solved), res)
	assert.Equal(t, Definitions(example.Horizontal), h)
	assert.Equal(t, Definitions(example.Vertical), v)
}
