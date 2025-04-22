package grid

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	"strings"
)

const (
	EmptyCell = '_'
	BlackCell = '#'
)

type Grid [][]rune

func New(width, height int) Grid {
	g := NewEmpty(width, height)
	g.drawBlackCells(2, 2)
	return g
}

func NewEmpty(width, height int) Grid {
	g := make(Grid, height)
	for i := range height {
		g[i] = make([]rune, width)
		for j := range g[i] {
			g[i][j] = EmptyCell
		}
	}
	return g
}

func (g Grid) drawBlackCells(maxLine, maxCol int) {
	countCol := func(j int) int {
		var count int
		for i := range g[j] {
			if g[i][j] == BlackCell {
				count++
			}
		}

		return count
	}

	for i := range g.Height() {
		inLine := rand.IntN(maxLine + 1)
		slog.Debug("drawing black cells in line", "line", i, "count", inLine)
		for count := 0; count < inLine; count++ {
			for _, j := range rand.Perm(g.Width()) {
				if countCol(j) < maxCol {
					slog.Debug("marking black cell", "line", i, "column", j)
					g[i][j] = BlackCell
					break
				}
			}
		}
	}
}

func NewRandom() Grid {
	randN := func(n int) int {
		return 6 + rand.IntN(n)
	}
	return New(randN(8), randN(8))
}

// ExampleGrid is the grid for the example in ./example.
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
var ExampleGrid = Grid{
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

func (g Grid) Clone() Grid {
	clone := make([][]rune, len(g))
	for i := range g {
		clone[i] = make([]rune, len(g[i]))
		copy(clone[i], g[i])
	}
	return clone
}

func (g Grid) Display() {
	maxLineSize := len(fmt.Sprintf("%d", g.Height()))
	leftPadding := func() { fmt.Print(strings.Repeat(" ", maxLineSize+1)) }

	// Won't go over 100 rows.
	leftPadding()
	for i := 1; i <= g.Width(); i++ {
		if i >= 10 {
			fmt.Printf("%d ", i/10)
		} else {
			fmt.Print("  ")
		}
	}
	fmt.Println()

	leftPadding()
	for i := 1; i <= g.Width(); i++ {
		fmt.Printf("%d ", i%10)
	}
	fmt.Println()

	for i, row := range g {
		padding := strings.Repeat(" ", maxLineSize-len(fmt.Sprintf("%d", i+1)))
		fmt.Printf("%s%d ", padding, i+1)

		for _, char := range row {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}
}

func (g Grid) Print() {
	for _, row := range g {
		for _, char := range row {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}
}

func (g Grid) Width() int {
	return len(g[0])
}

func (g Grid) Height() int {
	return len(g)
}

type Segment struct {
	Line, Start, Length int
}

func (g Grid) FindLineSegments(line int) []Segment {
	var res []Segment

	var start int
	for i, char := range g[line] {
		if char != BlackCell {
			start = i
			break
		}
	}

	for i := start; i < len(g[line]); i++ {
		if g[line][i] == BlackCell {
			if l := i - start; l > 1 {
				res = append(res, Segment{line, start, i - start})
			}
			start = i + 1
		}
	}

	if l := g.Width() - start; start < g.Width() && l > 1 {
		res = append(res, Segment{line, start, g.Width() - start})
	}

	return res
}

func (g Grid) FindAllLineSegments() []Segment {
	res := g.FindLineSegments(0)

	for i := 1; i < g.Height(); i++ {
		res = append(res, g.FindLineSegments(i)...)
	}

	return res
}

func (g Grid) FillLineSegment(line, column int, word string) string {
	previous := make([]rune, len(word))
	for j, c := range []rune(word) {
		previous[j] = g[line][column+j]
		g[line][column+j] = c
	}
	return string(previous)
}

func (g Grid) EmptyLineSegment(line, column, length int) {
	for j := range length {
		g[line][column+j] = EmptyCell
	}
}

func (g Grid) UnFillLineSegment(line, column int) {
	for j := column; j < g.Width() && g[line][j] != BlackCell; j++ {
		g[line][j] = EmptyCell
	}
}

func (g Grid) FillColumnSegment(line, column int, word string) string {
	previous := make([]rune, len(word))
	for i, c := range []rune(word) {
		previous[i] = g[line+i][column]
		g[line+i][column] = c
	}
	return string(previous)
}

func (g Grid) UnFillColumnSegment(line, column int) {
	for i := line; i < g.Height() && g[line][column] != BlackCell; i++ {
		g[i][column] = EmptyCell
	}
}

func (g Grid) PreviousBlackCellInColumn(line, column int) int {
	i := line - 1
	for ; i >= 0 && g[i][column] != BlackCell; i-- {
	}
	return i
}

func (g Grid) WordsInColumn(column int) []string {
	var concat []rune
	for i := 0; i < g.Height(); i++ {
		concat = append(concat, g[i][column])
	}

	return strings.Split(strings.Trim(string(concat), string(BlackCell)), string(BlackCell))
}

func (g Grid) CompletionState() string {
	all, filled := 0, 0
	for i := range g.Height() {
		for j := range g.Width() {
			switch g[i][j] {
			case BlackCell:
				continue
			case EmptyCell:
			default:
				filled++
			}
			all++
		}
	}

	return fmt.Sprintf("%.0f%%", float64(filled)/float64(all)*100)
}
