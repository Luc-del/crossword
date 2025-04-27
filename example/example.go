package example

import (
	"crossword/dictionary"
	"crossword/grid"
)

func Dictionary() dictionary.LengthOrdered {
	return dictionary.NewLengthOrdered("words-example.json")
}

// Grid is the grid for the example in ./example.
// ..  0    1    2    3    4    5    6    7    8    9
// A {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'}, 0
// B {'.', '.', '.', 'X', '.', '.', '.', '.', '.', '.'}, 1
// C {'.', '.', 'X', '.', '.', '.', '.', '.', 'X', '.'}, 2
// D {'.', '.', '.', '.', '.', '.', 'X', '.', '.', 'X'}, 3
// E {'.', '.', '.', 'X', '.', '.', '.', 'X', '.', '.'}, 4
// F {'.', '.', '.', '.', '.', 'X', '.', '.', '.', '.'}, 5
// G {'.', 'X', '.', '.', 'X', '.', '.', '.', '.', '.'}, 6
// H {'.', '.', '.', '.', '.', '.', 'X', '.', '.', 'X'}, 7
// I {'.', '.', '.', 'X', 'X', '.', '.', '.', '.', '.'}, 8
// J {'.', '.', 'X', '.', '.', '.', '.', '.', '.', '.'}, 9
var Grid = grid.Grid{
	{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
	{EmptyCell, EmptyCell, EmptyCell, BlackCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
	{EmptyCell, EmptyCell, BlackCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, BlackCell, EmptyCell},
	{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, BlackCell, EmptyCell, EmptyCell, BlackCell},
	{EmptyCell, EmptyCell, EmptyCell, BlackCell, EmptyCell, EmptyCell, EmptyCell, BlackCell, EmptyCell, EmptyCell},
	{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, BlackCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
	{EmptyCell, BlackCell, EmptyCell, EmptyCell, BlackCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
	{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, BlackCell, EmptyCell, EmptyCell, BlackCell},
	{EmptyCell, EmptyCell, EmptyCell, BlackCell, BlackCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
	{EmptyCell, EmptyCell, BlackCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
}

const (
	EmptyCell = grid.EmptyCell
	BlackCell = grid.BlackCell
)

var (
	Solved = [][]rune{
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

	Horizontal = [][]string{
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
	}

	Vertical = [][]string{
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
	}
)
