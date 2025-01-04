package grid

import (
	"fmt"
	"strings"
)

const (
	EmptyCell = '_'
	BlackCell = '#'
)

type Grid [][]rune

func NewGrid() Grid {
	// TODO handle size
	// TODO random X number
	// TODO random X location
	return ExampleGrid
}

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
	Start, Length int
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
				res = append(res, Segment{start, i - start})
			}
			start = i + 1
		}
	}

	if l := g.Width() - start; start < g.Width() && l > 1 {
		res = append(res, Segment{start, g.Width() - start})
	}

	return res
}

func (g Grid) FillLineSegment(line, column int, word string) {
	for j, c := range []rune(word) {
		g[line][column+j] = c
	}
}

func (g Grid) FillColumnSegment(line, column int, word string) {
	for i, c := range []rune(word) {
		g[line+i][column] = c
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
