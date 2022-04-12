package main

import (
	"GO_PROJ/handlers"
	"log"
	"net/http"
)

func main() {

	router := handlers.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
