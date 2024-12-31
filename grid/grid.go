package grid

import "fmt"

const (
	EmptyCell = '.'
	BlackCell = 'X'
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

func (g Grid) Clone() [][]rune {
	clone := make([][]rune, len(g))
	for i := range g {
		clone[i] = make([]rune, len(g[i]))
		copy(clone[i], g[i])
	}
	return clone
}

func (g Grid) Print() {
	for _, row := range g {
		for _, char := range row {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}
}

type segment struct {
	start, length int
}

//func (g Grid) FindSegments(i int) []segment {
//	var res []segment
//	start := -1
//
//	for i, char := range g[i] {
//		if char == '.' {
//			if start == -1 {
//				start = i
//			}
//		} else {
//			if start != -1 {
//				segments = append(segments, struct{ start, length int }{start, i - start})
//				start = -1
//			}
//		}
//	}
//	if start != -1 {
//		segments = append(segments, struct{ start, length int }{start, len(line) - start})
//	}
//	return segments
//}
