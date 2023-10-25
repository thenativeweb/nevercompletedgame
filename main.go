package main

import (
	"fmt"
	"net/http"

	"github.com/thenativeweb/ncg/api"
	"github.com/thenativeweb/ncg/environment"
	"github.com/thenativeweb/ncg/storage/inmemory"
)

func main() {
	port, err := environment.Port(3000)
	if err != nil {
		panic(err)
	}

	store := inmemory.New()

	router := api.GetRouter(store)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
