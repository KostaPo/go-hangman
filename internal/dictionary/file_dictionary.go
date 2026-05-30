package dictionary

import (
	"bufio"
	"hangman/assets"
	"hangman/internal/word"
	"math/rand"
	"strings"
)

type FileDictionary struct {
	words []string
}

func NewFileDictionary() *FileDictionary {

	words, err := readWordsFromFile()

	if err != nil {
		panic(err)
	}

	if len(words) == 0 {
		panic("список слов пуст")
	}

	return &FileDictionary{words: words}
}

func (d *FileDictionary) GetRandomWord() *word.Word {
	strWord := d.words[rand.Intn(len(d.words))]
	return word.NewWord(strWord)
}

func readWordsFromFile() ([]string, error) {
	var wordsFromFile []string

	scanner := bufio.NewScanner(strings.NewReader(assets.WordsFile))
	for scanner.Scan() {
		if word := scanner.Text(); word != "" {
			wordsFromFile = append(wordsFromFile, word)
		}
	}

	return wordsFromFile, scanner.Err()
}
