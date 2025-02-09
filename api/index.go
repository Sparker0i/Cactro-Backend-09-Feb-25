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

func init() {
	maxSize := 10
	if val, exists := os.LookupEnv("MAX_CACHE_SIZE"); exists {
		if parsed, err := strconv.Atoi(val); err == nil && parsed > 0 {
			maxSize = parsed
		} else {
			fmt.Printf("Invalid MAX_CACHE_SIZE value '%s'. Using default: %d\n", val, maxSize)
		}
	}
	fmt.Printf("Initializing cache with maximum size: %d\n", maxSize)

	cacheInstance := cache.New(maxSize)
	cacheHandler := handlers.New(cacheInstance)

	router = gin.Default()
	router.POST("/cache", cacheHandler.PostCacheHandler)
	router.GET("/cache/:key", cacheHandler.GetCacheHandler)
	router.DELETE("/cache/:key", cacheHandler.DeleteCacheHandler)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
