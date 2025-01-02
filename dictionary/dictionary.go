package dictionary

import (
	"crossword/utils"
	"encoding/json"
	"os"
	"regexp"
)

type Dictionary map[string]string // word:definition

func New(path string) Dictionary {
	d, err := loadWordsFromJSON(path)
	if err != nil {
		panic(err)
	}

	return d
}

func NewDefault() Dictionary {
	d, err := loadWordsFromJSON("dictionary/words-example.json")
	if err != nil {
		panic(err)
	}

	return d
}

func NewExample() Dictionary {
	d, err := loadWordsFromJSON("dictionary/words-example.json")
	if err != nil {
		panic(err)
	}

	return d
}

func (d *Dictionary) Remove(word string) {
	delete(*d, word)
}

func loadWordsFromJSON(filePath string) (Dictionary, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words Dictionary
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&words); err != nil {
		return nil, err
	}

	return words, nil
}

func (d Dictionary) ContainsMatch(regex string) (string, bool) {
	r := regexp.MustCompile(regex)
	for word := range d {
		if r.MatchString(utils.RemoveAccentString(word)) {
			return word, true
		}
	}
	return "", false
}
