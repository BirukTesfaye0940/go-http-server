package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}

	server := NewServer()

	log.Println("Server running on :8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
