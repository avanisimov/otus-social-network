package main

import (
	"log"
	"net/http"

	httpDelivery "github.com/avanisimov/otus-social-network/internal/http"
	"github.com/avanisimov/otus-social-network/internal/user"
	"github.com/avanisimov/otus-social-network/internal/db"
)

func main() {

	database, err := db.Connect("localhost", "postgres", "postgres", "social", 5432)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
		// Test connection
	if err := database.Ping(); err != nil {
		log.Fatal("failed to ping DB:", err)
	}

	log.Println("Connected to PostgreSQL")

	userRepo := user.NewRepository(database)
	userService := user.NewService(userRepo)

	router := httpDelivery.NewRouter(userService)

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", router)
}
