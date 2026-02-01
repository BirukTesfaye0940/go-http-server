package main

import (
	"fmt"
	"net/http"
	"os"

	"go-http-server/handlers"
	"go-http-server/middleware"
	"go-http-server/models"
	"go-http-server/storage"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 2. Auto-Migrate: This creates the "users" table automatically!
	db.AutoMigrate(&models.User{})
	return db, nil
}

func NewServer() *http.Server {
	db, err := initDB()
	if err != nil {
		panic("failed to connect database")
	}
	repo := storage.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.Users)

	handler := middleware.Recovery(middleware.Logging(middleware.Auth(mux)))

	return &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: handler,
	}
}
