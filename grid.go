package main

func NewGrid() [][]rune {
	// TODO handle size
	// TODO random X number
	// TODO random X location
	return exampleGrid
}

var exampleGrid = [][]rune{
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', 'X', '.', '.', '.', '.', '.', '.'},
	{'.', '.', 'X', '.', '.', '.', '.', '.', 'X', '.'},
	{'.', '.', '.', '.', '.', '.', 'X', '.', '.', 'X'},
	{'.', '.', '.', 'X', '.', '.', '.', 'X', '.', '.'},
	{'.', '.', '.', '.', '.', 'X', '.', '.', '.', '.'},
	{'.', 'X', '.', '.', 'X', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', 'X', '.', '.', 'X'},
	{'.', '.', '.', 'X', 'X', '.', '.', '.', '.', '.'},
	{'.', '.', 'X', '.', '.', '.', '.', '.', '.', '.'},
}