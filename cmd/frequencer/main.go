package main

import (
	"bufio"
	"encoding/json"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	wordIdx      = 0
	freqFilmIdx  = 8
	freqLivreIdx = 9
)

type wordData struct {
	Word string `json:"word"`
	Freq int    `json:"freq"`
}

func main() {
	file, err := os.Open("./dictionary/assets/raw/Lexique383.tsv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // omit first line

	freqOrdered := make([]wordData, 0, 50000)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "\t")

		word := values[wordIdx]
		if !isSimpleWord(word) {
			continue
		}

		current := wordData{
			Word: word,
			Freq: freq(values[freqFilmIdx], values[freqLivreIdx]),
		}

		switch {
		case len(freqOrdered) < cap(freqOrdered):
			freqOrdered = append(freqOrdered, current)
		case current.Freq > freqOrdered[len(freqOrdered)-1].Freq:
			freqOrdered[len(freqOrdered)-1] = current
		default:
			continue
		}

		sort.Slice(freqOrdered, func(i, j int) bool {
			return freqOrdered[i].Freq > freqOrdered[j].Freq
		})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	//debug, err := json.Marshal(freqOrdered)
	//if err != nil {
	//	panic(err)
	//}
	//if err := os.WriteFile("debug_frequence.json", debug, 0644); err != nil {
	//	panic(err)
	//}

	words := make(map[string]string, len(freqOrdered))
	for _, w := range freqOrdered {
		words[w.Word] = ""
	}

	data, err := json.Marshal(words)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("./dictionary/assets/frequenced.json", data, 0644); err != nil {
		panic(err)
	}
}

func freq(freqFilm, freqLivre string) int {
	ff, err := strconv.ParseFloat(freqFilm, 64)
	if err != nil {
		panic(err)
	}
	fl, err := strconv.ParseFloat(freqLivre, 64)
	if err != nil {
		panic(err)
	}
	return int(100 * (2*ff + fl) / 3)
}

func isSimpleWord(s string) bool {
	for _, r := range s {
		if r == ' ' || r == '"' || r == '\'' {
			return false
		}
	}
	strings.ContainsAny(s, " \"'\\,-")
	return len(s) > 1
}
