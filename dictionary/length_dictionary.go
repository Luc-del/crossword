package dictionary

import (
	"crossword/grid"
	"regexp"
)

type LengthOrdered map[int]map[string]string

func NewLengthOrdered(fileName string) LengthOrdered {
	d, err := LoadWordsFromJSON(fileName)
	if err != nil {
		panic(err)
	}

	ordered := make(LengthOrdered)
	for k, v := range d {
		if ordered[len(k)] == nil {
			ordered[len(k)] = make(map[string]string)
		}
		ordered[len(k)][k] = v
	}

	return ordered
}

func (d LengthOrdered) Remove(word string) {
	delete(d[len(word)], word)
}

func (d LengthOrdered) Pop(word string) (string, bool) {
	w, ok := d[len(word)][word]
	if ok {
		delete(d[len(word)], word)
	}
	return w, ok
}

func (d LengthOrdered) Add(word, def string) {
	d[len(word)][word] = def
}

// ContainsMatch works only info the regex is given without repetititors, for a known sized word
func (d LengthOrdered) ContainsMatch(regex string) (string, bool) {
	length := len(regex) - 2
	r := regexp.MustCompile(regex)
	for word := range d[length] {
		if r.MatchString(word) {
			return word, true
		}
	}
	return "", false
}

func (d LengthOrdered) ContainsMatchN(regex string, atLeast int) (string, int) {
	length := len(regex) - 2
	r := regexp.MustCompile(regex)

	var (
		count int
		match string
	)
	for word := range d[length] {
		if r.MatchString(word) {
			count++
			match = word
			if count == atLeast {
				return match, count
			}
		}
	}
	return match, count
}

func matchPattern(word string, pattern []rune) bool {
	for i, r := range pattern {
		if r != grid.EmptyCell && rune(word[i]) != r {
			return false
		}
	}
	return true
}

func (d LengthOrdered) ContainsPatternN(pattern []rune, atLeast int) (string, int) {
	var (
		count int
		match string
	)
	for word := range d[len(pattern)] {
		if matchPattern(word, pattern) {
			count++
			match = word
			if count == atLeast {
				return match, count
			}
		}
	}
	return match, count
}

func (d LengthOrdered) Registry(regex string) map[string]string {
	return d[len(regex)-2]
}
