package main

import (
	"go-htmx/internal/appmux"
	"go-htmx/internal/logger"
	"log"
	"net/http"
)

func main() {

	logger.InitLogger()

	mux := appmux.New()
	err := http.ListenAndServe("0.0.0.0:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
