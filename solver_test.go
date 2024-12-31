package main

import (
	"crossword/grid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildConstraint(t *testing.T) {
	g := grid.ExampleGrid.Clone()

	tester := func(t *testing.T, line, column int, expected string) {
		assert.Equal(t, expected, buildConstraint(g, '_', line, column))
	}

	t.Run("bounded by X", func(t *testing.T) {
		tester(t, 5, 3, "_**")
		tester(t, 6, 3, "._*")
		tester(t, 7, 3, ".._")
	})

	t.Run("start of grid", func(t *testing.T) {
		tester(t, 0, 1, "_*****")
		tester(t, 1, 1, "._****")
		tester(t, 5, 1, "....._")
	})

	t.Run("end of grid", func(t *testing.T) {
		tester(t, 1, 1, "._****")
	})

	t.Run("start & end of grid", func(t *testing.T) {
		tester(t, 0, 0, "_*********")
		tester(t, 1, 0, "._********")
		tester(t, 9, 0, "........._")
	})
}
