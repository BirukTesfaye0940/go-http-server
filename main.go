package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Load configuration
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}

	server := NewServer()

	// 2. Create a channel to listen for shutdown signals from the OS
	// (Signals like Ctrl+C or Docker's termination signal)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// 3. Start the server in a "goroutine" (background thread)
	// We do this because ListenAndServe blocks execution.
	go func() {
		log.Printf("Server running on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %s\n", err)
		}
	}()

	// 4. The code "hangs" here until we receive a signal in the 'stop' channel
	<-stop
	log.Println("Shutting down server gracefully...")

	// 5. Create a deadline for the shutdown process (e.g., 5 seconds)
	// This prevents the server from hanging forever if a request is stuck.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 6. Tell the server to shut down using the timeout context
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server successfully stopped. Goodbye!")
}
