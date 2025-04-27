package main

import (
	"crossword/grid"
	"fmt"
	"strings"

	"github.com/phpdave11/gofpdf"
)

const cellSize = 6.0

func main() {
	g := grid.Grid([][]rune{
		{'O', 'M', 'N', 'I', 'S', 'C', 'I', 'E', 'N', 'T'},
		{'C', 'H', 'A', 'M', 'A', 'I', 'L', 'L', 'E', '#'},
		{'C', '#', 'N', 'A', '#', 'B', 'L', 'A', 'N', 'C'},
		{'U', 'L', 'T', 'R', 'A', 'L', 'E', 'G', 'E', 'R'},
		{'P', 'I', 'U', '#', 'R', 'E', 'T', 'U', 'S', 'A'},
		{'E', '#', 'A', 'L', 'E', 'R', 'T', 'E', '#', 'P'},
		{'R', 'O', 'T', 'I', 'N', '#', 'R', 'U', 'N', 'E'},
		{'#', 'L', 'I', 'M', 'A', 'C', 'E', 'S', '#', 'T'},
		{'O', 'L', 'E', 'A', 'C', 'E', '#', 'E', 'S', 'T'},
		{'M', 'E', 'N', 'N', 'E', 'U', 'S', '#', 'T', 'E'},
	})

	blackCells := g.BlackCellPositions()

	definitions := map[string]string{
		"omniscient": "Sait tout",
		"chamaille":  "Petite bagarre sans gravité",
		"na":         "Sodium",
		"blanc":      "Couleur par défaut",
		"ultraleger": "Plus léger que léger",
		"piu":        "Encore plus, en musique",
		"retusa":     "Sorte de bonzaï Ficus",
		"alerte":     "Prêt à réagir à tout moment",
		"rotin":      "Fibre végétale tressée",
		"rune":       "Symbole ancien gravé",
		"limaces":    "Gastéropodes sans coquille",
		"oleace":     "Famille botanique de l'olivier",
		"est":        "Direction du soleil levant",
		"menneus":    "Genre d'araignée aranéomorphe",
		"te":         "Pronom personnel",

		"occuper":    "Prendre place",
		"om":         "Club de foot de Provence",
		"mh":         "Abréviation scientifique de masse hydrique",
		"li":         "Symbole du lithium",
		"olle":       "Commune française d'Eure et Loire",
		"nantuatien": "Habitant de Nantua",
		"imar":       "Dérivé marin ou rare",
		"liman":      "Embouchure d'un fleuve",
		"sa":         "La sienne",
		"arenace":    "Roche détritique, qualifie un sol sablonneux",
		"cibler":     "Viser précisément",
		"ceu":        "Ancien pluriel de \"ciel\" en occitan",
		"illettre":   "Ne sachant ni lire ni écrire",
		"elagueuse":  "Machine à couper les branches",
		"nenes":      "Terme familier pour \"seins\"",
		"st":         "Abréviation de \"saint\"",
		"crapette":   "Petit jeu de cartes à deux",
	}

	var horizontals = make([][]string, g.Height())
	var hdef = make([][]string, g.Height())
	for i := range g.Height() {
		for _, w := range g.WordsInLine(i) {
			if len(w) > 1 {
				w = strings.ToLower(w)
				horizontals[i] = append(horizontals[i], w)
				hdef[i] = append(hdef[i], definitions[w])
			}
		}
	}

	var verticals = make([][]string, g.Width())
	var vdef = make([][]string, g.Width())
	for j := range g.Width() {
		for _, w := range g.WordsInColumn(j) {
			if len(w) > 1 {
				w = strings.ToLower(w)
				verticals[j] = append(verticals[j], w)
				vdef[j] = append(vdef[j], definitions[w])
			}
		}
	}

	fmt.Println("horizontals", horizontals)
	fmt.Println("verticals", verticals)
	fmt.Println("hdef", hdef)
	fmt.Println("vdef", vdef)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("Roboto", "", "C:\\Users\\lucde\\go\\src\\crossword\\cmd\\pdf\\Roboto.ttf")
	pdf.SetFont("Roboto", "", 12)

	drawGrid(pdf, 10, 10, g.Height(), g.Width(), cellSize, blackCells, nil)
	pdf.SetY(10 + float64(g.Height())*cellSize)

	pdf.Ln(5)
	pdf.Cell(0, 10, "Définitions horizontales:")
	pdf.Ln(7)
	pdf.MultiCell(0, 7, format(hdef), "", "", false)

	pdf.Ln(5)
	pdf.Cell(0, 10, "Définitions Verticales:")
	pdf.Ln(7)
	pdf.MultiCell(0, 7, format(vdef), "", "", false)

	// Print solutions
	pdf.AddPage()
	words := gridIndices(g)
	drawGrid(pdf, 10, 10, g.Height(), g.Width(), cellSize, blackCells, words)
	pdf.SetY(10 + float64(g.Height())*cellSize)

	pdf.Ln(10)
	pdf.Cell(0, 10, "Solutions Horizontales:")
	pdf.Ln(7)
	pdf.MultiCell(0, 7, format(horizontals), "", "", false)

	pdf.Ln(5)
	pdf.Cell(0, 10, "Solutions Verticales:")
	pdf.Ln(7)
	pdf.MultiCell(0, 7, format(verticals), "", "", false)

	err := pdf.OutputFileAndClose("./pdf/grille_mots_croises.pdf")
	if err != nil {
		panic(err)
	}
}

// Fonction pour dessiner une grille
func drawGrid(pdf *gofpdf.Fpdf, startX, startY float64, rows, cols int, cellSize float64, blackCells map[[2]int]bool, letters map[[2]int]string) {
	pdf.SetXY(startX, startY)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x := startX + float64(j)*cellSize
			y := startY + float64(i)*cellSize

			if blackCells[[2]int{i, j}] {
				pdf.SetFillColor(0, 0, 0) // noir
				pdf.Rect(x, y, cellSize, cellSize, "F")
			} else {
				pdf.Rect(x, y, cellSize, cellSize, "")
				if letters != nil {
					if l, ok := letters[[2]int{i, j}]; ok {
						pdf.Text(x+cellSize/4, y+3*cellSize/4, l)
					}
				}
			}
		}
	}
}

func format(def [][]string) string {
	var res string
	for i := range def {
		res += fmt.Sprintf("- %d. %s\n", i+1, strings.Join(def[i], ". "))
	}

	return res
}

func gridIndices(g grid.Grid) map[[2]int]string {
	res := map[[2]int]string{}
	for i := range g.Height() {
		for j := range g.Width() {
			if c := g[i][j]; c != grid.BlackCell {
				res[[2]int{i, j}] = string(c)
			}
		}
	}
	return res
}
