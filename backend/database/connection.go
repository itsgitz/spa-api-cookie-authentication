package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DatabaseConnection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

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
