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
	// Imho wird hier uuid.New() auf Funktionalität getestet. Der Test kann imho nur fehlschlagen wenn uuid.New() eine
	// ID liefert, die bereits vorhanden ist.
	// Das liegt aber nicht im Verantwortungsbereich der App und muss damit imho auch nicht getestet werden.
	t.Run("Returns nil if the game does not exist yet.", func(t *testing.T) {
		store := inmemory.New()
		g := game.Start()

		err := store.Create(g)
		assert.NoError(t, err)
	})

	// Siehe oben.
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
	// Wieso existiert das Game nicht? Da steht doch eindeutig game.Start()?
	// Gemeint ist wahrscheinlich, dass das Game im Store nicht bekannt ist. Das reflektiert der Test-Name aber nicht.
	// Um den Test zu verstehen, muss ich internes Wissen über die Funktionsweise der App besitzen. Nicht gut.
	t.Run("Returns an error if the game does not exist.", func(t *testing.T) {
		store := inmemory.New()
		g := game.Start()

		err := store.Update(g)
		assert.Error(t, err)
	})

	// "Updates the game." sollte wohl eher "An Update of the game should be reflected in the store." heißen.
	// Beim ersten Lesen war mir nicht klar was der Test bezwecken soll.
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
