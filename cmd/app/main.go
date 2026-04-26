package main

import (
	"log"
	"net/http"

	httpDelivery "github.com/avanisimov/otus-social-network/internal/http"
	"github.com/avanisimov/otus-social-network/internal/user"
	"github.com/avanisimov/otus-social-network/internal/db"
	cfg "github.com/avanisimov/otus-social-network/internal/config"
)

func main() {

	config := cfg.Load()

	database, err := db.Connect(
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
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

	router := httpDelivery.NewRouter(userService, database)

	log.Println("Listening on :" + config.AppPort)
	http.ListenAndServe(":" + config.AppPort, router)
}

