package game

import (
	"unicode"

	"hangman/internal/dictionary"
	"hangman/internal/models"
	"hangman/internal/player"
	"hangman/internal/renderer"
)

type Game struct {
	gDictionary dictionary.Dictionary
	gRender     renderer.Renderer
}

func NewGame(render renderer.Renderer, dictionary dictionary.Dictionary) *Game {
	return &Game{
		gDictionary: dictionary,
		gRender:     render,
	}
}

func (g *Game) Start() {

	// Стартовое меню один раз
	choice := g.gRender.DrawMenu(g.newMenuState())
	if choice == 1 {
		return
	}

	for {
		state := g.getNewDataState()

		// Игровой цикл
		g.playGame(state)

		//Меню результата (победа / поражение)
		choice = g.gRender.DrawMenu(*state)
		if choice == 1 {
			return
		}
	}
}

func (g *Game) newMenuState() models.DataState {
	return models.DataState{
		GameState: models.InMenuState,
		Player:    *player.NewPlayer(),
	}
}

func (g *Game) playGame(state *models.DataState) {
	// Отрисовываем начальное состояние до первого ввода
	g.gRender.Draw(*state)

	for state.GameState == models.InGameState {
		// Читаем буквы в raw-режиме (без Enter)
		input := g.gRender.ReadLetter()
		state = g.processInput(state, input)
		g.gRender.Draw(*state)
	}
}

func (g *Game) getNewDataState() *models.DataState {
	p := player.NewPlayer()
	w := g.gDictionary.GetRandomWord()

	return &models.DataState{
		GameState:   models.InGameState,
		Word:        *w,
		Player:      *p,
		AlreadyUsed: false,
		WrongInput:  false,
	}
}

func (g *Game) processInput(state *models.DataState, input rune) *models.DataState {

	if !unicode.IsLetter(input) {
		state.WrongInput = true
		state.AlreadyUsed = false
		return state
	}

	letter := unicode.ToLower(input)

	if !state.Player.AttemptUniq(letter) {
		state.AlreadyUsed = true
		state.WrongInput = false
		return state
	}

	state.WrongInput = false
	state.AlreadyUsed = false

	ok := state.Word.Guess(letter)
	if !ok {
		state.Player.LoseHealth()
	}

	if state.Word.IsComplete() {
		state.GameState = models.InWinState
		return state
	}

	if !state.Player.IsAlive() {
		state.GameState = models.InLoseState
	}

	return state
}
