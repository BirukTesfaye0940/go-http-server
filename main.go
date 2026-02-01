package main

import (
	"log"
)

func main() {
	server := NewServer()

	log.Println("Server running on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
