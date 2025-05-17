package v4

import (
	"crossword/example"
	"crossword/grid"
	"crossword/utils/logger"
	"fmt"
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

	if !assert.Equal(t, grid.Grid(example.Solved), res) {
		fmt.Println("expected")
		grid.Grid(example.Solved).Print()

		fmt.Println("got")
		res.Print()
	}
	assert.Equal(t, Definitions(example.Horizontal), h)
	assert.Equal(t, Definitions(example.Vertical), v)
}
