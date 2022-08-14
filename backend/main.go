package main

import (
	"go-api/database"
	"go-api/server"
)

func init() {
	database.InitDatabase()
}

func main() {
	server.Run()
}
