package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/cache"
	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/handlers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// The init function sets up the cache and routes.
func init() {
	// Determine the maximum cache size from the environment variable.
	maxSize := 10 // default value
	if val, exists := os.LookupEnv("MAX_CACHE_SIZE"); exists {
		if parsed, err := strconv.Atoi(val); err == nil && parsed > 0 {
			maxSize = parsed
		} else {
			fmt.Printf("Invalid MAX_CACHE_SIZE value '%s'. Using default: %d\n", val, maxSize)
		}
	}
	fmt.Printf("Initializing cache with maximum size: %d\n", maxSize)

	// Initialize the cache and HTTP handlers.
	cacheInstance := cache.New(maxSize)
	cacheHandler := handlers.NewCacheHandler(cacheInstance)

	// Set up the Gin router with the defined routes.
	router = gin.Default()
	router.POST("/cache", cacheHandler.PostCacheHandler)
	router.GET("/cache/:key", cacheHandler.GetCacheHandler)
	router.DELETE("/cache/:key", cacheHandler.DeleteCacheHandler)
}

// Handler is the exported function that Vercel will invoke for every HTTP request.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Use Gin's ServeHTTP to handle the request.
	router.ServeHTTP(w, r)
}
