package main

import (
	"hangman/internal/dictionary"
	"hangman/internal/game"
	"hangman/internal/renderer"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			println("Ошибка получения словаря:", r.(string))
			return
		}
	}()

	rend := renderer.NewConsoleRenderer()
	dict := dictionary.NewFileDictionary()

	g := game.NewGame(rend, dict)
	g.Start()
}
