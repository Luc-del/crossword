package solver

import (
	"crossword/dictionary"
	"crossword/grid"
	"crossword/utils/logger"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	logger.Configure(slog.LevelDebug)
}

func TestFindCandidate(t *testing.T) {

}

func TestVerifyCandidate(t *testing.T) {
	t.Run("empty dictionary", func(t *testing.T) {
		s := New(dictionary.Dictionary{}, grid.ExampleGrid.Clone())
		_, ok := s.verifyCandidate("COQUELICOT", 0, 0)
		assert.False(t, ok)
	})

	t.Run("no words matching", func(t *testing.T) {
		s := New(dictionary.Dictionary{"ANTICONSTITUTIONNELEMENT": ""}, grid.ExampleGrid.Clone())
		_, ok := s.verifyCandidate("COQUELICOT", 0, 0)
		assert.False(t, ok)
	})

	t.Run("one word missing", func(t *testing.T) {
		d := dictionary.NewExample()
		delete(d, "OUILLE")

		s := New(d, grid.ExampleGrid.Clone())
		_, ok := s.verifyCandidate("COQUELICOT", 0, 0)
		assert.False(t, ok)
	})

	t.Run("matching words", func(t *testing.T) {
		s := New(dictionary.NewExample(), grid.ExampleGrid.Clone())
		_, ok := s.verifyCandidate("COQUELICOT", 0, 0)
		assert.True(t, ok)
	})
}

func TestBuildColumnConstraint(t *testing.T) {
	tester := func(t *testing.T, line, column int, expected string) {
		g := grid.ExampleGrid.Clone()
		s := Solver{g: g}
		assert.Equal(t, expected, s.buildColumnConstraint('¤', line, column))
	}

	t.Run("bounded by X", func(t *testing.T) {
		tester(t, 5, 3, "^¤..$")
		tester(t, 6, 3, "^.¤.$")
		tester(t, 7, 3, "^..¤$")
	})

	t.Run("start of grid", func(t *testing.T) {
		tester(t, 0, 1, "^¤.....$")
		tester(t, 1, 1, "^.¤....$")
		tester(t, 5, 1, "^.....¤$")
	})

	t.Run("end of grid", func(t *testing.T) {
		tester(t, 1, 1, "^.¤....$")
	})

	t.Run("start & end of grid", func(t *testing.T) {
		tester(t, 0, 0, "^¤.........$")
		tester(t, 1, 0, "^.¤........$")
		tester(t, 9, 0, "^.........¤$")
	})

	t.Run("no constraint on single character", func(t *testing.T) {
		tester(t, 0, 3, "")
		tester(t, 9, 4, "")
	})

	t.Run("column already filled is not a constraint", func(t *testing.T) {
		g := grid.ExampleGrid.Clone()
		g.FillColumnSegment(0, 1, "OUILLE")
		s := Solver{g: g}
		assert.Equal(t, "", s.buildColumnConstraint('U', 1, 1))
	})
}

func TestBuildLineConstraint(t *testing.T) {
	t.Run("filling segment", func(t *testing.T) {
		s := Solver{g: [][]rune{{'_', '_', '_', '#', '_', '_', '_', '_'}}}
		assert.Equal(t, "^...$", s.buildLineSegmentConstraint(0, 0, 3))
		assert.Equal(t, "^....$", s.buildLineSegmentConstraint(0, 4, 4))
	})

	t.Run("segment already filled isn't a constraint", func(t *testing.T) {
		s := Solver{g: [][]rune{{'a', 'b', 'c', '#', 'd', 'e', 'f', 'g'}}}
		assert.Equal(t, "", s.buildLineSegmentConstraint(0, 0, 3))
		assert.Equal(t, "", s.buildLineSegmentConstraint(0, 4, 4))
	})
}

func TestSolve(t *testing.T) {
	s := New(dictionary.NewExample(), grid.ExampleGrid.Clone())
	s.fillLineSegment(4, 4, "NEO")
	s.fillLineSegment(4, 8, "OR")

	expected := [][]rune{
		{'C', 'O', 'Q', 'U', 'E', 'L', 'I', 'C', 'O', 'T'},
		{'O', 'U', 'I', '#', 'R', 'A', 'C', 'I', 'N', 'E'},
		{'C', 'I', '#', 'O', 'R', 'T', 'I', 'E', '#', 'K'},
		{'C', 'L', 'I', 'M', 'A', 'T', '#', 'L', 'P', '#'},
		{'I', 'L', 'S', '#', 'N', 'E', 'O', '#', 'O', 'R'},
		{'N', 'E', 'A', 'N', 'T', '#', 'H', 'A', 'I', 'E'},
		{'E', '#', 'T', 'E', '#', 'S', 'E', 'R', 'R', 'E'},
		{'L', 'A', 'I', 'T', 'U', 'E', '#', 'B', 'E', '#'},
		{'L', 'I', 'S', '#', '#', 'V', 'I', 'R', 'A', 'L'},
		{'E', 'L', '#', 'E', 'L', 'E', 'V', 'E', 'U', 'R'},
	}

	h, v, res := s.Solve()
	assert.Equal(t, grid.Grid(expected), res)
	assert.Empty(t, s.d)

	assert.Equal(t,
		Definitions([][]string{
			{"Met de la couleur dans les champs."},
			{"Bonne réponse.", "Pousse sous terre."},
			{"Démonstratif.", "Qui s'y frotte s'y pique."},
			{"Fait la pluie et le beau temps.", "33 tours."},
			{"Après vous.", "Héros de Matrix.", "Élément de valeur."},
			{"Vraiment rien.", "Barrière végétale."},
			{"Règle à dessin.", "Jardin d'hiver."},
			{"Romaine au potager.", "La Belgique sur le web."},
			{"Tigre ou Martagon.", "Donc contagieux."},
			{"Article étranger.", "Producteur de viande."},
		}),
		h,
	)

	assert.Equal(t,
		Definitions([][]string{
			{"Rouge à points noirs."},
			{"Variété d'aïe.", "Planté dans le potager."},
			{"À 140 c'est génial.", "Pastel des teinturiers."},
			{"Onze marseillais.", "Après impôts."},
			{"Vagabond."},
			{"Pièce de bois.", "Monte au printemps."},
			{"Pas ailleurs.", "Formule d'appel.", "Quatre à Rome."},
			{"Bleu l'été.", "Fruiter au jardin."},
			{"Vecteur de rumeur.", "Cultivé au jardin."},
			{"Bois précieux.", "Crie sous les bois.", "Symbole pour le lawrencium."},
		}),
		v,
	)
}
