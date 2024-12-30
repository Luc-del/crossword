package main

import (
	"encoding/json"
	"os"
)

var dictionary []Word

func init() {
	var err error
	dictionary, err = loadWordsFromJSON("words-example.json")
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
	err = decoder.Decode(&words)
	if err != nil {
		return nil, err
	}
	return words, nil
}
