package main

import (
	"crossword/dictionary"
	"crossword/grid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyCandidate(t *testing.T) {
	t.Run("empty dictionary", func(t *testing.T) {
		s := solver{
			g: grid.ExampleGrid.Clone(),
			d: dictionary.Dictionary{},
		}

		assert.False(t, s.verifyCandidate("coquelicot", 0, 0))
	})

	t.Run("no words matching", func(t *testing.T) {
		s := solver{
			g: grid.ExampleGrid.Clone(),
			d: dictionary.Dictionary{"anticonstitutionnelement": ""},
		}

		assert.False(t, s.verifyCandidate("coquelicot", 0, 0))
	})

	t.Run("one word missing", func(t *testing.T) {
		d := dictionary.NewExample()
		delete(d, "ouille")

		s := solver{
			g: grid.ExampleGrid.Clone(),
			d: d,
		}

		assert.False(t, s.verifyCandidate("coquelicot", 0, 0))
	})

	t.Run("matching words", func(t *testing.T) {
		s := solver{
			g: grid.ExampleGrid.Clone(),
			d: dictionary.NewExample(),
		}

		assert.True(t, s.verifyCandidate("coquelicot", 0, 0))
	})
}

func TestBuildConstraint(t *testing.T) {
	tester := func(t *testing.T, line, column int, expected string) {
		g := grid.ExampleGrid.Clone()
		assert.Equal(t, expected, solver{g: g}.buildConstraint('¤', line, column))
	}

	t.Run("bounded by X", func(t *testing.T) {
		tester(t, 5, 3, "^¤.{2}$")
		tester(t, 6, 3, "^_¤.{1}$")
		tester(t, 7, 3, "^__¤$")
	})

	t.Run("start of grid", func(t *testing.T) {
		tester(t, 0, 1, "^¤.{5}$")
		tester(t, 1, 1, "^_¤.{4}$")
		tester(t, 5, 1, "^_____¤$")
	})

	t.Run("end of grid", func(t *testing.T) {
		tester(t, 1, 1, "^_¤.{4}$")
	})

	t.Run("start & end of grid", func(t *testing.T) {
		tester(t, 0, 0, "^¤.{9}$")
		tester(t, 1, 0, "^_¤.{8}$")
		tester(t, 9, 0, "^_________¤$")
	})

	t.Run("no constraint on single character", func(t *testing.T) {
		tester(t, 0, 3, "")
		tester(t, 9, 4, "")
	})
}
