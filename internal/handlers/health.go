package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler provides endpoints for liveness and readiness probes.
type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Liveness handles GET /live and returns a simple JSON indicating that the application is alive.
func (h *HealthHandler) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}

// Readiness handles GET /ready and returns a simple JSON indicating that the application is ready.
func (h *HealthHandler) Readiness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}
