package game

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

var ErrIncorrectGuess = errors.New("incorrect guess")
var ErrGameAlreadyCompleted = errors.New("game already completed")

type Game struct {
	ID    uuid.UUID
	level int
}

func Start() *Game {
	return &Game{
		ID:    uuid.New(),
		level: 0,
	}
}

func (g *Game) MakeGuess(guess string) error {
	if g.IsCompleted() {
		return ErrGameAlreadyCompleted
	}

	answer := levels[g.level].Answer

	if !strings.EqualFold(guess, answer) {
		return ErrIncorrectGuess
	}

	g.level++

	return nil
}

func (g *Game) Level() int {
	return g.level + 1
}

func (g *Game) Question() string {
	return levels[g.level].Question
}

func (g *Game) IsCompleted() bool {
	return g.level >= len(levels)
}
