package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("./dictionary/assets/raw/unimontreal-dem.jsonl")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	words, err := readJSONL(file)
	if err != nil {
		panic(err)
	}

	res, err := json.Marshal(words)
	if err != nil {
		panic(err)
	}

	outputName := fmt.Sprintf("./dictionary/assets/extracted.%d.json", time.Now().Unix())
	if err := os.WriteFile(outputName, res, 0644); err != nil {
		panic(err)
	}
}

func readJSONL(file io.Reader) (map[string]string, error) {
	output := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := gjson.Get(scanner.Text(), "M").String()

		if !filter(word) {
			continue
		}

		output[word] = ""
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return output, nil
}

func filter(word string) bool {
	return !strings.ContainsAny(word, " -")
}
