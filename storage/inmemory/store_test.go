package inmemory_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/ncg/game"
	"github.com/thenativeweb/ncg/storage/inmemory"
)

func TestNew(t *testing.T) {
	t.Run("Returns a new store.", func(t *testing.T) {
		store := inmemory.New()

		assert.NotNil(t, store)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Returns nil if the game does not exist yet.", func(t *testing.T) {
		store := inmemory.New()
		g := game.Start()

		err := store.Create(g)
		assert.NoError(t, err)
	})

	t.Run("Returns an error if the game already exists.", func(t *testing.T) {
		store := inmemory.New()
		g := game.Start()

		err := store.Create(g)
		assert.NoError(t, err)

		err = store.Create(g)
		assert.Error(t, err)
	})

	t.Run("Stores the game.", func(t *testing.T) {
		store := inmemory.New()
		g := game.Start()

		err := store.Create(g)
		assert.NoError(t, err)

		storedGame, err := store.ReadByID(g.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, g, storedGame)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Returns an error if the game does not exist.", func(t *testing.T) {
		store := inmemory.New()
		g := game.Start()

		err := store.Update(g)
		assert.Error(t, err)
	})

	t.Run("Updates the game.", func(t *testing.T) {
		store := inmemory.New()
		g := game.Start()

		err := store.Create(g)
		assert.NoError(t, err)

		err = g.MakeGuess("A")
		assert.NoError(t, err)

		err = store.Update(g)
		assert.NoError(t, err)

		storedGame, err := store.ReadByID(g.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, g, storedGame)
	})
}

func TestReadByID(t *testing.T) {
	t.Run("Returns the game if it exists.", func(t *testing.T) {
		store := inmemory.New()
		g := game.Start()

		err := store.Create(g)
		assert.NoError(t, err)

		storedGame, err := store.ReadByID(g.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, g, storedGame)
	})

	t.Run("Returns an error if the game does not exist.", func(t *testing.T) {
		store := inmemory.New()

		_, err := store.ReadByID("foo")
		assert.Error(t, err)
	})
}
