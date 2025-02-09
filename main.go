package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/config"
	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/server"
)

func main() {
	// Load configuration (e.g., MAX_CACHE_SIZE from environment variables)
	cfg := config.LoadConfig()
	fmt.Printf("Initializing cache with maximum size: %d\n", cfg.MaxCacheSize)

	// Create the Gin router by injecting the configuration and dependencies.
	router := server.NewRouter(cfg)

	// Create the HTTP server using the router.
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start the server in a separate goroutine.
	go func() {
		log.Println("Server is running on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %s", err)
		}
	}()

	// Set up a channel to listen for interrupt or termination signals.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit // Block until a signal is received.
	log.Println("Shutting down server...")

	// Create a deadline for the graceful shutdown (e.g., 5 seconds).
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server.
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	log.Println("Server exiting")
}
