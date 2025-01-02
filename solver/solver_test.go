package solver

import (
	"crossword/dictionary"
	"crossword/grid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCandidate(t *testing.T) {

}

func TestVerifyCandidate(t *testing.T) {
	t.Run("empty dictionary", func(t *testing.T) {
		s := Solver{
			g: grid.ExampleGrid.Clone(),
			d: dictionary.Dictionary{},
		}

		assert.False(t, s.verifyCandidate("coquelicot", 0, 0))
	})

	t.Run("no words matching", func(t *testing.T) {
		s := Solver{
			g: grid.ExampleGrid.Clone(),
			d: dictionary.Dictionary{"anticonstitutionnelement": ""},
		}

		assert.False(t, s.verifyCandidate("coquelicot", 0, 0))
	})

	t.Run("one word missing", func(t *testing.T) {
		d := dictionary.NewExample()
		delete(d, "ouille")

		s := Solver{
			g: grid.ExampleGrid.Clone(),
			d: d,
		}

		assert.False(t, s.verifyCandidate("coquelicot", 0, 0))
	})

	t.Run("matching words", func(t *testing.T) {
		s := Solver{
			g: grid.ExampleGrid.Clone(),
			d: dictionary.NewExample(),
		}

		assert.True(t, s.verifyCandidate("coquelicot", 0, 0))
	})
}

func TestBuildColumnConstraint(t *testing.T) {
	tester := func(t *testing.T, line, column int, expected string) {
		g := grid.ExampleGrid.Clone()
		assert.Equal(t, expected, Solver{g: g}.buildColumnConstraint('¤', line, column))
	}

	t.Run("bounded by X", func(t *testing.T) {
		tester(t, 5, 3, "^¤.{2}$")
		tester(t, 6, 3, "^.¤.{1}$")
		tester(t, 7, 3, "^..¤$")
	})

	t.Run("start of grid", func(t *testing.T) {
		tester(t, 0, 1, "^¤.{5}$")
		tester(t, 1, 1, "^.¤.{4}$")
		tester(t, 5, 1, "^.....¤$")
	})

	t.Run("end of grid", func(t *testing.T) {
		tester(t, 1, 1, "^.¤.{4}$")
	})

	t.Run("start & end of grid", func(t *testing.T) {
		tester(t, 0, 0, "^¤.{9}$")
		tester(t, 1, 0, "^.¤.{8}$")
		tester(t, 9, 0, "^.........¤$")
	})

	t.Run("no constraint on single character", func(t *testing.T) {
		tester(t, 0, 3, "")
		tester(t, 9, 4, "")
	})
}

func TestBuildLineConstraint(t *testing.T) {
	s := Solver{g: [][]rune{{'_', '_', '_', '#', '_', '_', '_', '_'}}}
	assert.Equal(t, "^...$", s.buildLineSegmentConstraint(0, 0, 3))
	assert.Equal(t, "^....$", s.buildLineSegmentConstraint(0, 4, 4))
}

func TestSolve(t *testing.T) {
	g := grid.ExampleGrid.Clone()
	g.FillLineSegment(4, 4, "neo")
	g.FillLineSegment(4, 8, "or")

	s := Solver{
		g: g,
		d: dictionary.NewExample(),
	}

	expected := [][]rune{
		{'c', 'o', 'q', 'u', 'e', 'l', 'i', 'c', 'o', 't'},
		{'o', 'u', 'i', '#', 'r', 'a', 'c', 'i', 'n', 'e'},
		{'c', 'i', '#', 'o', 'r', 't', 'i', 'e', '#', 'k'},
		{'c', 'l', 'i', 'm', 'a', 't', '#', 'l', 'p', '#'},
		{'i', 'l', 's', '#', 'n', 'e', 'o', '#', 'o', 'r'},
		{'n', 'é', 'a', 'n', 't', '#', 'h', 'a', 'i', 'e'},
		{'e', '#', 't', 'e', '#', 's', 'e', 'r', 'r', 'e'},
		{'l', 'a', 'i', 't', 'u', 'e', '#', 'b', 'e', '#'},
		{'l', 'i', 's', '#', '#', 'v', 'i', 'r', 'a', 'l'},
		{'e', 'l', '#', 'é', 'l', 'e', 'v', 'e', 'u', 'r'},
	}

	_, _, res := s.Solve()
	assert.Equal(t, grid.Grid(expected), res)
}
