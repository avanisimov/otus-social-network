package main

import (
	"log"
	"net/http"

	httpDelivery "github.com/avanisimov/otus-social-network/internal/http"
	"github.com/avanisimov/otus-social-network/internal/user"
)

func main() {

	userRepo := user.NewRepository()
	userService := user.NewService(userRepo)

	router := httpDelivery.NewRouter(userService)

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", router)
}
