package server

import (
	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/cache"
	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/config"
	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/handlers"
	"github.com/gin-gonic/gin"
)

// NewRouter creates and returns a configured Gin router using the provided configuration.
func NewRouter(cfg *config.Config) *gin.Engine {
	// Initialize the cache with the configured maximum size.
	cacheInstance := cache.New(cfg.MaxCacheSize)
	cacheHandler := handlers.New(cacheInstance)

	// Create a Gin router with default middleware (Logger, Recovery).
	router := gin.Default()

	// Register endpoints.
	router.POST("/cache", cacheHandler.PostCacheHandler)
	router.GET("/cache/:key", cacheHandler.GetCacheHandler)
	router.DELETE("/cache/:key", cacheHandler.DeleteCacheHandler)

	return router
}
