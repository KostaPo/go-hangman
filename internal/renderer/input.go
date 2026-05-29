package renderer

import (
	"os"

	"golang.org/x/term"
)

type MenuKey int

const (
	MenuUnknown MenuKey = iota
	MenuLeft
	MenuRight
	MenuEnter
)

type Input struct {
	oldState *term.State
	raw      bool
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) EnableRaw() error {
	if i.raw {
		return nil
	}

	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	i.oldState = state
	i.raw = true

	return nil
}

func (i *Input) DisableRaw() {
	if !i.raw {
		return
	}

	_ = term.Restore(int(os.Stdin.Fd()), i.oldState)
	i.raw = false
}

func (i *Input) ReadKey() []byte {
	buf := make([]byte, 8)

	n, _ := os.Stdin.Read(buf)

	return buf[:n]
}

func (i *Input) ReadMenuKey() MenuKey {
	_ = i.EnableRaw()
	defer i.DisableRaw()

	for {
		key := i.ReadKey()

		switch {
		case IsEnter(key):
			return MenuEnter

		case IsLeft(key):
			return MenuLeft

		case IsRight(key):
			return MenuRight

		case len(key) == 1 && key[0] == '1':
			return MenuEnter

		case len(key) == 1 && key[0] == '2':
			return MenuRight
		}
	}
}

func IsEnter(b []byte) bool {
	return len(b) == 1 && (b[0] == 13 || b[0] == 10)
}

func IsLeft(b []byte) bool {
	return len(b) >= 3 &&
		b[0] == 27 &&
		b[1] == '[' &&
		b[2] == 'D'
}

func IsRight(b []byte) bool {
	return len(b) >= 3 &&
		b[0] == 27 &&
		b[1] == '[' &&
		b[2] == 'C'
}
