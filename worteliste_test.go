package main

import (
	"testing"
)

func TestOpenFile(t *testing.T) {
	var wl WordlistI = &Wordlist{filename: WORD_LIST_FILE_NAME}
	wordList, err := wl.Loadlist()
	if len(wordList) == 0 {
		t.Errorf("The list must contain words. Size of Array is: %d\n Err msg: %s", len(wordList), err)
	}

}
func TestWordlistLength(t *testing.T) {
	var wl WordlistI = &Wordlist{filename: WORD_LIST_FILE_NAME}
	wordList, err := wl.Loadlist()
	if len(wordList) != WORD_LIST_SIZE {
		t.Errorf("The list must contain 4999 words. Size of Array is: %d\n Err msg: %s", len(wordList), err)
	}

}

func TestWordlistLastWord(t *testing.T) {
	var wl WordlistI = &Wordlist{filename: WORD_LIST_FILE_NAME}
	wordList, err := wl.Loadlist()
	if wordList[WORD_LIST_SIZE-1] != "üblicheres" {
		t.Errorf("Last word must be üblicheres %s\n Err msg: %s", wordList[WORD_LIST_SIZE-1], err)
	}

}
