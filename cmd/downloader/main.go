package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type dict struct {
	M string `json:"M"`
}

func main() {
	//out, err := os.Create("output.json")
	//if err != nil {
	//	panic(err)
	//}
	//defer out.Close()

	dictFile, err := os.Create("unimontreal-dem.jsonl")
	if err != nil {
		panic(err)
	}
	defer dictFile.Close()

	resp, err := http.Get("http://rali.iro.umontreal.ca/LVF_DEM/DEM.jsonl")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("bad status: %s", resp.Status))
	}

	if _, err := io.Copy(dictFile, resp.Body); err != nil {
		panic(err)
	}

	/*
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()

			var value dict
			if err := json.Unmarshal([]byte(line), &value); err != nil {
				panic(err)
			}

			if !keep(value.M) {
				continue
			}

			if _, err := fmt.Fprintln(out, value.M); err != nil {
				panic(err)
			}

		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	*/
}
