package startgame

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thenativeweb/ncg/game"
	"github.com/thenativeweb/ncg/storage"
)

func Handle(store storage.Store) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		currentGame := game.Start()

		err := store.Create(currentGame)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody := ResponseBody{
			GameID:   currentGame.ID.String(),
			Level:    currentGame.Level(),
			Question: currentGame.Question(),
		}

		responseBytes, err := json.Marshal(responseBody)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(responseBytes)
	}
}
