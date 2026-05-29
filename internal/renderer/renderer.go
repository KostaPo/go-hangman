package renderer

import "hangman/internal/models"

type Renderer interface {
	Draw(state models.DataState)
	DrawMenu(state models.DataState) int
	ReadLetter() rune
}
