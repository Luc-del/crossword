package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"regexp"
)

type Dictionary []Word

func (d *Dictionary) Remove(word string) {
	dict := *d
	for i, w := range *d {
		if w.Text == word {
			dict[i] = dict[len(dict)-1]
			*d = dict[:len(dict)-1]
			return
		}
	}
}

var words Dictionary

func init() {
	var err error
	words, err = loadWordsFromJSON("words-example.json")
	if err != nil {
		panic(err)
	}
}

type Word struct {
	Text       string
	Definition string
}

func loadWordsFromJSON(filePath string) ([]Word, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []Word
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&words); err != nil {
		return nil, err
	}

	return words, nil
}

func (d Dictionary) containsMatch(regex string) bool {
	r := regexp.MustCompile(regex)
	for i := range d {
		if r.MatchString(d[i].Text) {
			return true
		}
	}
	return false
}

func (d Dictionary) randomWord() Word {
	return d[rand.Intn(len(d))]
}
