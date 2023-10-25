package ping

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Handle() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		status := http.StatusOK

		w.WriteHeader(status)
		w.Write([]byte(http.StatusText(status)))
	}
}
