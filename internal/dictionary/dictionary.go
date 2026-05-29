package dictionary

import "hangman/internal/word"

type Dictionary interface {
	GetRandomWord() *word.Word
}
