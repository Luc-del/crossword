package main

import (
	"encoding/json"
	"math/rand"
	"os"
)

var dictionary []Word

func init() {
	var err error
	dictionary, err = loadWordsFromJSON("words-sample.json")
	if err != nil {
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
	err = decoder.Decode(&words)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func getRandomWord(dictionary []Word, usedWords map[string]bool) Word {
	for {
		idx := rand.Intn(len(dictionary))

		if _, ok := usedWords[dictionary[idx].Text]; !ok {
			return dictionary[idx]
		}
	}
}
