package dictionary

import (
	"bufio"
	"hangman/internal/word"
	"math/rand"
	"os"
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

	file, err := os.Open("assets/words.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var wordsFromFile []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if word := scanner.Text(); word != "" {
			wordsFromFile = append(wordsFromFile, word)
		}
	}

	return wordsFromFile, scanner.Err()
}
