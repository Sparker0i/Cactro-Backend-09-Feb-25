package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Cache struct {
	mu      sync.RWMutex
	data    map[string]string
	maxSize int
}

func NewCache(maxSize int) *Cache {
	return &Cache{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
}

func (c *Cache) Set(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.data[key]; exists {
		c.data[key] = value
		return nil
	}

	if len(c.data) >= c.maxSize {
		return fmt.Errorf("cache is full: maximum of %d item size reached", c.maxSize)
	}

	c.data[key] = value
	return nil
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, exists := c.data[key]
	return value, exists
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func main() {
	maxSize := 10
	if val, exists := os.LookupEnv("MAX_CACHE_SIZE"); exists {
		if parsed, err := strconv.Atoi(val); err == nil && parsed > 0 {
			maxSize = parsed
		} else {
			fmt.Printf("Invalid MAX_CACHE_SIZE value %s; Using default %d\n", val, maxSize)
		}
	}

	router := gin.Default()
	cache := NewCache(maxSize)

	router.POST("/cache", func(c *gin.Context) {
		var req struct {
			Key   string `json:"key" binding:"required"`
			Value string `json:"value" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := cache.Set(req.Key, req.Value); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "stored successfully"})
	})

	router.GET("/cache/:key", func(c *gin.Context) {
		key := c.Param("key")

		value, exists := cache.Get(key)

		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("key not found %s", key)})
		}

		c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
	})

	router.DELETE("/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		cache.Delete(key)
		c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
	})

	router.Run(":8080")
}
