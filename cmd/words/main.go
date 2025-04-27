package main

import (
	"crossword/grid"
	"fmt"
)

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

	fmt.Println("Horizontals:")
	for i := range g.Height() {
		fmt.Printf("%d: ", i)
		for _, w := range g.WordsInLine(i) {
			if len(w) > 1 {
				fmt.Printf("%s ", w)
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Verticals:")
	for j := range g.Width() {
		fmt.Printf("%d: ", j)
		for _, w := range g.WordsInColumn(j) {
			if len(w) > 1 {
				fmt.Printf("%s ", w)
			}
		}
		fmt.Println()
	}
}

/*
                     1
   1 2 3 4 5 6 7 8 9 0
 1 O M N I S C I E N T
 2 C H A M A I L L E #
 3 C # N A # B L A N C
 4 U L T R A L E G E R
 5 P I U # R E T U S A
 6 E # A L E R T E # P
 7 R O T I N # R U N E
 8 # L I M A C E S # T
 9 O L E A C E # E S T
10 M E N N E U S # T E

*/
