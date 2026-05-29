package renderer

import (
	"fmt"
	"hangman/assets"
	"hangman/internal/models"
	"strings"
)

type BottomPanel struct{}

func (p *BottomPanel) Render(s models.DataState) string {
	var sb strings.Builder

	sb.WriteString(assets.Divider + "\n\n")

	switch s.GameState {
	case models.InGameState:
		sb.WriteString(p.renderInput(s))
	}

	return sb.String()
}

func (p *BottomPanel) renderInput(s models.DataState) string {
	var sb strings.Builder

	if s.AlreadyUsed {
		sb.WriteString(fmt.Sprintf("  %s\n\n", assets.MsgAlreadyUsed))
	} else if s.WrongInput {
		sb.WriteString(fmt.Sprintf("  %s\n\n", assets.MsgWrongInput))
	} else {
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("  %s", assets.MsgPrompt))

	return sb.String()
}

func (p *BottomPanel) RenderMenu(
	s models.DataState,
	selected int,
) string {

	var sb strings.Builder

	sb.WriteString(assets.Divider + "\n\n")

	switch s.GameState {

	case models.InMenuState:
		sb.WriteString("  Добро пожаловать!\n\n")

	case models.InWinState:
		sb.WriteString(fmt.Sprintf(
			"  %s\n\n",
			assets.MsgWin,
		))

	case models.InLoseState:
		sb.WriteString(fmt.Sprintf(
			"  "+assets.MsgLose+"\n\n",
			s.Word.DisplayOriginal(),
		))
	}

	options := []string{
		assets.MenuNewGame,
		assets.MenuExit,
	}

	sb.WriteString("  ")

	for i, option := range options {

		if i == selected {
			sb.WriteString(
				fmt.Sprintf(
					"\x1b[7m %s \x1b[0m  ",
					option,
				),
			)
		} else {
			sb.WriteString(
				fmt.Sprintf(" %s   ", option),
			)
		}
	}

	sb.WriteString("\n")
	sb.WriteString("  \x1b[90m(←/→, Enter)\x1b[0m\n")

	return sb.String()
}
