package main

import (
	"go-htmx/internal/appmux"
	"log"
	"net/http"
)

func main() {

	mux := appmux.New()

	err := http.ListenAndServe("0.0.0.0:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
