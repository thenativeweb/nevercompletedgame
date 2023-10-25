package game_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/ncg/game"
)

func TestStart(t *testing.T) {
	t.Run("Starts a game in level 1.", func(t *testing.T) {
		g := game.Start()

		assert.NotNil(t, g)
		assert.Equal(t, 1, g.Level())
	})

	t.Run("Starts a game with a unique ID.", func(t *testing.T) {
		g1 := game.Start()
		g2 := game.Start()

		assert.NotNil(t, g1)
		assert.NotNil(t, g2)
		assert.NotEqual(t, g1.ID, g2.ID)
	})
}

func TestMakeGuess(t *testing.T) {
	t.Run("Returns nil if the guess is correct.", func(t *testing.T) {
		g := game.Start()

		err := g.MakeGuess(game.AnswerForLevel(1))
		assert.NoError(t, err)
	})

	t.Run("Accepts guesses in any case.", func(t *testing.T) {
		answer := game.AnswerForLevel(1)
		answerUppercase := strings.ToUpper(answer)
		answerLowercase := strings.ToLower(answer)

		g := game.Start()
		err := g.MakeGuess(answerUppercase)
		assert.NoError(t, err)

		g = game.Start()
		err = g.MakeGuess(answerLowercase)
		assert.NoError(t, err)
	})

	t.Run("Returns an error if the guess is incorrect.", func(t *testing.T) {
		g := game.Start()

		err := g.MakeGuess("this should return an error because this guess is incorrect")
		assert.Error(t, err)
	})

	t.Run("Returns an error if the game is already completed.", func(t *testing.T) {
		g := game.Start()

		for _, answer := range game.Answers() {
			g.MakeGuess(answer)
		}

		assert.True(t, g.IsCompleted())

		err := g.MakeGuess("this should return an error because the game is already completed")
		assert.Error(t, err)
	})

	t.Run("Proceeds to the next level if the guess is correct.", func(t *testing.T) {
		g := game.Start()

		err := g.MakeGuess(game.AnswerForLevel(1))
		assert.NoError(t, err)
		assert.Equal(t, 2, g.Level())
	})
}

func TestQuestion(t *testing.T) {
	t.Run("Returns the question for the current level.", func(t *testing.T) {
		g := game.Start()

		question := g.Question()
		assert.Equal(t, game.QuestionForLevel(1), question)

		g.MakeGuess(game.AnswerForLevel(1))

		question = g.Question()
		assert.Equal(t, game.QuestionForLevel(2), question)
	})
}

func TestIsCompleted(t *testing.T) {
	t.Run("Returns false if the game is not completed.", func(t *testing.T) {
		g := game.Start()

		assert.False(t, g.IsCompleted())
	})

	t.Run("Returns true if the game is completed.", func(t *testing.T) {
		g := game.Start()

		for _, answer := range game.Answers() {
			g.MakeGuess(answer)
		}

		assert.True(t, g.IsCompleted())
	})
}
