package player

import "unicode"

const MaxHealth = 5

type Player struct {
	Health   int
	attempts map[rune]struct{}
}

func NewPlayer() *Player {
	return &Player{
		Health:   MaxHealth,
		attempts: make(map[rune]struct{}),
	}
}

func (p *Player) GetAttempts() map[rune]struct{} {
	return p.attempts
}

func (p *Player) IsAlive() bool {
	return p.Health > 0
}

func (p *Player) LoseHealth() {
	p.Health--
}

func (p *Player) AttemptUniq(letter rune) bool {
	letter = unicode.ToLower(letter)
	if _, exists := p.attempts[letter]; exists {
		return false
	}
	p.attempts[letter] = struct{}{}
	return true
}
