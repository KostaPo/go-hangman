package models

import (
	"hangman/internal/player"
	"hangman/internal/word"
)

type DataState struct {
	GameState   GameState
	Word        word.Word
	Player      player.Player
	AlreadyUsed bool
	WrongInput  bool
}
