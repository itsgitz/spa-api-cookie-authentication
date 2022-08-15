package app

import (
	"go-api/database"
	"go-api/server"

	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	database.InitDatabase()
	server.RunServer()
}
