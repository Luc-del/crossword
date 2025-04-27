package matcher

import "crossword/grid"

func Pattern(word string, pattern []rune) bool {
	for i, r := range pattern {
		if r != grid.EmptyCell && rune(word[i]) != r {
			return false
		}
	}
	return true
}
