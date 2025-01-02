package main

import (
	"crossword/dictionary"
	"crossword/grid"
)

func main() {
	g := grid.ExampleGrid.Clone()
	g.FillLineSegment(4, 4, "neo")
	g.FillLineSegment(4, 8, "or")
	d := dictionary.NewExample()
	//d.Remove("neo")
	//d.Remove("or")
	s := solver{
		g: g,
		d: d,
	}
	g.Print()

	defer s.g.Print()
	s.solve()
	//horizontals, verticals, filledGrid := s.solve()
	//
	//fmt.Println("Grille initiale :")
	//s.g.Print()
	//
	//fmt.Println("\nSolution complète :")
	//filledGrid.Print()
	//
	//fmt.Println("\nDéfinitions horizontales :")
	//for k, def := range horizontals {
	//	fmt.Printf("Ligne %d: %s\n", k, def)
	//}
	//
	//fmt.Println("\nDéfinitions verticales :")
	//for k, def := range verticals {
	//	fmt.Printf("Colonne %d: %s\n", k, def)
	//}
}
