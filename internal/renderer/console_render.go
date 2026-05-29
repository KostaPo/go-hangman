package renderer

import (
	"fmt"
	"hangman/internal/models"
	"os"
)

type ConsoleRenderer struct {
	topPanel    *TopPanel
	bottomPanel *BottomPanel
	input       *Input
}

func NewConsoleRenderer() *ConsoleRenderer {
	return &ConsoleRenderer{
		topPanel:    &TopPanel{},
		bottomPanel: &BottomPanel{},
		input:       NewInput(),
	}
}

// render — полная перерисовка экрана.
func (r *ConsoleRenderer) render(content string) {
	r.input.DisableRaw()

	// полный reset терминала
	fmt.Fprint(os.Stdout, "\033[?25l") // hide cursor
	fmt.Fprint(os.Stdout, "\033[H")    // cursor home
	fmt.Fprint(os.Stdout, "\033[2J")   // clear screen
	fmt.Fprint(os.Stdout, "\033[3J")   // clear scrollback

	fmt.Fprint(os.Stdout, content)

	fmt.Fprint(os.Stdout, "\033[?25h") // show cursor
}

func (r *ConsoleRenderer) Draw(state models.DataState) {
	content :=
		r.topPanel.Render(state) +
			r.bottomPanel.Render(state)

	r.render(content)
}

func (r *ConsoleRenderer) DrawMenu(state models.DataState) int {
	sel := 0

	for {
		content :=
			r.topPanel.Render(state) +
				r.bottomPanel.RenderMenu(state, sel)

		r.render(content)

		key := r.input.ReadMenuKey()

		switch key {
		case MenuLeft:
			if sel > 0 {
				sel--
			}

		case MenuRight:
			if sel < 1 {
				sel++
			}

		case MenuEnter:
			return sel
		}
	}
}

func (r *ConsoleRenderer) ReadLetter() rune {
	_ = r.input.EnableRaw()

	for {
		key := r.input.ReadKey()

		if len(key) == 0 {
			continue
		}

		// Ctrl+C
		if key[0] == 3 {
			r.input.DisableRaw()
			return 0
		}

		runes := []rune(string(key))

		if len(runes) > 0 {
			r.input.DisableRaw()
			return runes[0]
		}
	}
}
