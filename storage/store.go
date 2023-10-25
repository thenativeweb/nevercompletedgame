package storage

import "github.com/thenativeweb/ncg/game"

type Store interface {
	Create(game *game.Game) error
	Update(game *game.Game) error

	ReadByID(gameID string) (*game.Game, error)
}
