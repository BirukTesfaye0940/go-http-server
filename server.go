package main

import (
	"net/http"
	"time"

	"go-http-server/handlers"
	"go-http-server/storage"
)

func NewServer() *http.Server {
	repo := storage.NewUserRepository()
	userHandler := handlers.NewUserHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.Users)

	return &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
