package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go-http-server/handlers"
	"go-http-server/middleware"
	"go-http-server/models"
	"go-http-server/storage"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initDB handles the database connection with a retry mechanism to wait for Postgres to be ready
func initDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var db *gorm.DB
	var err error

	// Retry loop: Attempt to connect 5 times with a 2-second delay between each
	for i := 1; i <= 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to the database!")
			break
		}

		log.Printf("Attempt %d: Database not ready, retrying in 2s... (%v)", i, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after retries: %w", err)
	}

	// Auto-Migrate: This creates the "users" table automatically!
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}

func NewServer() *http.Server {
	db, err := initDB()
	if err != nil {
		// Using log.Fatal is cleaner than panic for startup errors
		log.Fatalf("Critical error during server initialization: %v", err)
	}

	repo := storage.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.Users)

	handler := middleware.Recovery(middleware.Logging(middleware.Auth(mux)))

	// Default to 8080 if PORT env is missing
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
}
