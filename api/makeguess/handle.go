package makeguess

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thenativeweb/ncg/storage"
)

func Handle(store storage.Store) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		requestBytes, err := io.ReadAll(io.LimitReader(request.Body, 4096))
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		currentGame, err := store.ReadByID(requestBody.GameID)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		err = currentGame.MakeGuess(requestBody.Guess)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		err = store.Update(currentGame)
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
