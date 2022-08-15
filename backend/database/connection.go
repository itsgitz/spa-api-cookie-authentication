package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func DatabaseConnection() (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, err
}
