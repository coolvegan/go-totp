package main

import (
	"encoding/json"
	"os"
)

const (
	WORD_LIST_SIZE      = 4999
	WORD_LIST_FILE_NAME = "./wortliste.json"
)

type WordlistI interface {
	Loadlist() ([]string, error)
}

type Wordlist struct {
	filename string
}

func (w *Wordlist) Loadlist() ([]string, error) {
	var result []string
	wl, err := os.ReadFile(w.filename)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(wl, &result)

	return result, nil
}
