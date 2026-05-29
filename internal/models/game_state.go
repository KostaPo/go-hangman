package models

type GameState string

const (
	InMenuState GameState = "menu"
	InGameState GameState = "game"
	InWinState  GameState = "win"
	InLoseState GameState = "lose"
)
