package main

import (
	"encoding/json"
	"os"
	"strings"
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

	okWords := make(map[string]string, 10000)
	var count int
	for k := range words {
		if isSimpleWord(k) {
			okWords[strings.ToLower(k)] = words[k]
			count++
			if count == 10000 {
				break
			}
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

func isSimpleWord(s string) bool {
	for _, r := range s {
		if r == ' ' || r == '"' || r == '\'' {
			return false
		}
	}
	strings.ContainsAny(s, " \"'\\,-")
	return len(s) > 0
}
