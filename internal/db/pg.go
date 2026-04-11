package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(host, user, password, dbName string, port int) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbName, port,
	)

	return sql.Open("postgres", dsn)
}