package database

import (
	"context"
	"log"
	"time"
)

func InitDatabase() {
	RunMigration()
}

func RunMigration() {
	db, err := DatabaseConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Millisecond)
	defer cancel()

	query := `
	CREATE TABLE IF NOT EXISTS users(
		id serial PRIMARY KEY,
		username VARCHAR(256) UNIQUE NOT NULL,
		password VARCHAR(256) NOT NULL
	)
	`
	result, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	_ = rows

	log.Println("Database table migrated!")
}
