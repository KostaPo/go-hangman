package renderer

import (
	"fmt"
	"hangman/assets"
	"hangman/internal/models"
	"hangman/internal/player"
	"strings"
)

// TopPanel отвечает за верхнюю зону экрана.
type TopPanel struct{}

func (p *TopPanel) Render(s models.DataState) string {
	switch s.GameState {
	case models.InMenuState:
		return p.renderMenu()
	default:
		return p.renderGame(s)
	}
}

func (p *TopPanel) renderMenu() string {
	leftLines := strings.Split(assets.HangmanStages[0], "\n")

	rightLines := []string{
		"",
		" H A N G M A N",
		"",
		" консольная игра",
		"",
		"",
		"",
	}

	for len(rightLines) < len(leftLines) {
		rightLines = append(rightLines, "")
	}

	var sb strings.Builder
	for i, left := range leftLines {
		sb.WriteString(fmt.Sprintf("%-9s   %s\n", left, rightLines[i]))
	}
	return sb.String()
}

func (p *TopPanel) renderGame(s models.DataState) string {
	var sb strings.Builder
	sb.WriteString(p.renderColumns(s))
	sb.WriteString("\n")
	sb.WriteString(p.renderWord(s))
	sb.WriteString("\n")
	return sb.String()
}

func (p *TopPanel) renderColumns(s models.DataState) string {
	// Индекс стадии считаем от здоровья игрока
	lost := player.MaxHealth - s.Player.Health
	if lost < 0 {
		lost = 0
	}
	if lost >= len(assets.HangmanStages) {
		lost = len(assets.HangmanStages) - 1
	}

	leftLines := strings.Split(assets.HangmanStages[lost], "\n")
	rightLines := p.buildRightColumn(s)

	for len(rightLines) < len(leftLines) {
		rightLines = append(rightLines, "")
	}

	var sb strings.Builder
	for i, left := range leftLines {
		sb.WriteString(fmt.Sprintf("%-9s   %s\n", left, rightLines[i]))
	}
	return sb.String()
}

func (p *TopPanel) buildRightColumn(s models.DataState) []string {
	return []string{
		assets.MsgUsed,
		p.formatLetters(s.Player.GetAttempts()),
		"",
		"",
		fmt.Sprintf(assets.MsgLivesLeft, s.Player.Health),
		"",
		"",
	}
}

func (p *TopPanel) renderWord(s models.DataState) string {
	return fmt.Sprintf("  %s\n", s.Word.DisplayMask())
}

func (p *TopPanel) formatLetters(letters map[rune]struct{}) string {
	if len(letters) == 0 {
		return "—"
	}
	parts := make([]string, 0, len(letters))
	for r := range letters {
		parts = append(parts, string(r))
	}
	return strings.Join(parts, " ")
}
