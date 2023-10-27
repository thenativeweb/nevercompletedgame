package inmemory

import (
	"errors"

	"github.com/thenativeweb/ncg/game"
)

var ErrGameAlreadyExists = errors.New("game already exists")
var ErrGameNotFound = errors.New("game not found")

type Store struct {
	games map[string]*game.Game
}

func New() *Store {
	return &Store{
		games: make(map[string]*game.Game),
	}
}

func (s *Store) Create(game *game.Game) error {
	if _, ok := s.games[game.ID.String()]; ok {
		return ErrGameAlreadyExists
	}

	s.games[game.ID.String()] = game

	return nil
}

func (s *Store) Update(game *game.Game) error {
	if _, ok := s.games[game.ID.String()]; !ok {
		return ErrGameNotFound
	}

	s.games[game.ID.String()] = game

	return nil
}

func (s *Store) ReadByID(gameID string) (*game.Game, error) {
	// Variable 'game' collides with imported package name ;)
	game, ok := s.games[gameID]
	if !ok {
		return nil, ErrGameNotFound
	}

	return game, nil
}
