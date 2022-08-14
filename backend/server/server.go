package server

import (
	"go-api/handler"
	"log"
	"net/http"
	"time"
)

const (
	port = ":8000"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/login", handler.Login)

	server := serverConfiguration(mux)
	log.Println("Start API server on port", port)

	// Run HTTP Server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func serverConfiguration(mx *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:              port,
		Handler:           mx,
		ReadTimeout:       60 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
