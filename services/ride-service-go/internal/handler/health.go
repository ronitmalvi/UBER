package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ronitmalvi/UBER/ride-service-go/internal/config"
)

func Health(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
			"app":    cfg.AppName,
		})
	}
}