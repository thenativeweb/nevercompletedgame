package question

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thenativeweb/ncg/storage"
)

func Handle(store storage.Store) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		gameID := request.URL.Query().Get("gameId")
		currentGame, err := store.ReadByID(gameID)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody := ResponseBody{
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
