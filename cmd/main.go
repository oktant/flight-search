package main

import (
	"net/http"
	"github.com/srinivasvinay/flight-search/handlers"
	"log"
)

func main() {

	http.HandleFunc("/search", handlers.GetFlights)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
