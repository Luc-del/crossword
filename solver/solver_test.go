package solver

import (
	"crossword/dictionary"
	"crossword/grid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	d := dictionary.NewExample()
	d["tes"] = "several possibilities until the end, make a choice"

	h, v, res := Solve(d, grid.ExampleGrid.Clone())

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

	expectedHorizontal := Definitions([][]string{
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
	})

	expectedVertical := Definitions([][]string{
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
	})

	// Two possibilities left
	if res[2][9] == 'S' {
		expected[2][9] = 'S'
		expectedVertical[9][0] = "several possibilities until the end, make a choice"
	}

	assert.Equal(t, grid.Grid(expected), res)
	assert.Equal(t, expectedHorizontal, h)
	assert.Equal(t, expectedVertical, v)
}
