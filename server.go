package main

import (
	"net/http"

	"go-http-server/handlers"
	"go-http-server/middleware"
	"go-http-server/storage"
)

func NewServer() *http.Server {
	repo := storage.NewUserRepository()
	userHandler := handlers.NewUserHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.Users)

	handler := middleware.Recovery(middleware.Logging(middleware.Auth(mux)))

	return &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
}
