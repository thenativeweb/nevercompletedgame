package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/thenativeweb/ncg/api/makeguess"
	"github.com/thenativeweb/ncg/api/ping"
	"github.com/thenativeweb/ncg/api/question"
	"github.com/thenativeweb/ncg/api/startgame"
	"github.com/thenativeweb/ncg/storage"
)

func GetRouter(store storage.Store) *httprouter.Router {
	router := httprouter.New()

	// Technical supporting routes
	router.GET("/ping", ping.Handle())

	// Commands
	router.POST("/start-game", startgame.Handle(store))
	router.POST("/make-guess", makeguess.Handle(store))

	// Queries
	router.GET("/question", question.Handle(store))

	return router
}
