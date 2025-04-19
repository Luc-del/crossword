package main

import (
	"encoding/json"
	"os"
)

func main() {
	file, err := os.Open("./dictionary/assets/words.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var words map[string]string
	if err := json.NewDecoder(file).Decode(&words); err != nil {
		panic(err)
	}

	for k := range words {
		delete(words, k)
		if len(words) == 10000 {
			break
		}
	}

	data, err := json.Marshal(words)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("./dictionary/assets/words-shortened.json", data, 0644); err != nil {
		panic(err)
	}
}
