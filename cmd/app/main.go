package main

import (
	"log"
	"net/http"

	httpDelivery "github.com/avanisimov/otus-social-network/internal/http"
)

func main() {
	router := httpDelivery.NewRouter()

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", router)
}
