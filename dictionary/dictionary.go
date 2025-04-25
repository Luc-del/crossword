package dictionary

import (
	"embed"
	"encoding/json"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

//go:embed assets/*.json
var fs embed.FS

type Dictionary map[string]string // word:definition

func New(fileName string) Dictionary {
	d, err := loadWordsFromJSON(fileName)
	if err != nil {
		panic(err)
	}

	return d
}

func NewExample() Dictionary {
	d, err := loadWordsFromJSON("words-example.json")
	if err != nil {
		panic(err)
	}

	return d
}

func loadWordsFromJSON(fileName string) (Dictionary, error) {
	file, err := fs.Open("assets/" + fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words Dictionary
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&words); err != nil {
		return nil, err
	}

	formatted := make(Dictionary, len(words))
	for w, def := range words {
		formatted[strings.ToUpper(removeAccentString(w))] = def
	}

	return formatted, nil
}

func (d *Dictionary) Remove(word string) {
	delete(*d, word)
}

func (d *Dictionary) Pop(word string) (string, bool) {
	w, ok := (*d)[word]
	if ok {
		delete(*d, word)
	}
	return w, ok
}

func (d *Dictionary) Add(word, def string) {
	(*d)[word] = def
}

// ContainsMatch finds a match in the Dictionary.
func (d Dictionary) ContainsMatch(regex string) (string, bool) {
	r := regexp.MustCompile(regex)
	for word := range d {
		if r.MatchString(word) {
			return word, true
		}
	}
	return "", false
}

// ContainsMatchN finds if there is at least n matches in the Dictionary.
// It returns a match and the cardinality up to n.
func (d Dictionary) ContainsMatchN(regex string, atLeast int) (string, int) {
	r := regexp.MustCompile(regex)

	var (
		count int
		match string
	)
	for word := range d {
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

func (d Dictionary) Registry(string) map[string]string {
	return d
}

func removeAccentString(s string) string {
	res := make([]rune, utf8.RuneCountInString(s))
	for i, c := range []rune(s) {
		res[i] = removeAccent(c)
	}

	return string(res)
}

func removeAccent(r rune) rune {
	for _, char := range norm.NFD.String(string(r)) { // Normalize the string to decomposed form (NFD)
		if !unicode.Is(unicode.Mn, char) { // Mn: Non-spacing marks (accents)
			return char
		}
	}
	return r
}
