package main

import (
	"net/http"
	"time"

	"go-http-server/handlers"
	"go-http-server/storage"
	"log"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 1. Code here runs BEFORE the handler
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// 2. Pass the request to the real handler
		next.ServeHTTP(w, r)

		// 3. Code here runs AFTER the handler finishes
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func NewServer() *http.Server {
	repo := storage.NewUserRepository()
	userHandler := handlers.NewUserHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.Users)

	var handler http.Handler = mux
	handler = LoggingMiddleware(handler)

	return &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
