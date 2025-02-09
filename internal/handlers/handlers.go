package handlers

import (
	"fmt"
	"net/http"

	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/cache"
	"github.com/gin-gonic/gin"
)

type CacheHandler struct {
	Cache *cache.Cache
}

func NewCacheHandler(c *cache.Cache) *CacheHandler {
	return &CacheHandler{Cache: c}
}

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

func (h *CacheHandler) GetCacheHandler(c *gin.Context) {
	key := c.Param("key")

	value, exists := h.Cache.Get(key)

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("key not found %s", key)})
	}

	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}

func (h *CacheHandler) DeleteCacheHandler(c *gin.Context) {
	key := c.Param("key")
	h.Cache.Delete(key)
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
