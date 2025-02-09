package handlers

import (
	"net/http"

	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/cache"
	"github.com/gin-gonic/gin"
)

// CacheHandler holds the cache instance to be used by HTTP handlers.
type CacheHandler struct {
	Cache *cache.Cache
}

// NewCacheHandler creates a new CacheHandler.
func New(c *cache.Cache) *CacheHandler {
	return &CacheHandler{Cache: c}
}

// PostCacheHandler handles POST /cache to store a key-value pair.
func (h *CacheHandler) PostCacheHandler(c *gin.Context) {
	var req struct {
		Key   string `json:"key" binding:"required"`
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Cache.Set(req.Key, req.Value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "stored successfully"})
}

// GetCacheHandler handles GET /cache/:key to retrieve a stored value.
func (h *CacheHandler) GetCacheHandler(c *gin.Context) {
	key := c.Param("key")
	value, exists := h.Cache.Get(key)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "key not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}

// DeleteCacheHandler handles DELETE /cache/:key to remove a key from the cache.
func (h *CacheHandler) DeleteCacheHandler(c *gin.Context) {
	key := c.Param("key")
	h.Cache.Delete(key)
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
