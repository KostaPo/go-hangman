package word

import (
	"strings"
	"unicode"
)

type Word struct {
	target string
	hits   map[rune]bool
}

func NewWord(s string) *Word {
	s = strings.ToLower(s)
	return &Word{
		target: s,
		hits:   make(map[rune]bool),
	}
}

func (w *Word) Guess(letter rune) bool {
	letter = unicode.ToLower(letter)
	if strings.ContainsRune(w.target, letter) {
		w.hits[letter] = true
		return true
	}
	return false
}

func (w *Word) DisplayOriginal() string {
	return w.target
}

func (w *Word) DisplayMask() string {
	var display strings.Builder
	for _, ch := range w.target {
		if w.hits[ch] {
			display.WriteRune(ch)
		} else {
			display.WriteRune('_')
		}
	}
	return display.String()
}

func (w *Word) IsComplete() bool {
	for _, ch := range w.target {
		if !w.hits[ch] {
			return false
		}
	}
	return true
}
