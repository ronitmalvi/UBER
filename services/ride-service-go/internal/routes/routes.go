package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/ronitmalvi/UBER/ride-service-go/internal/config"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/handler"
)

func Register(router *gin.Engine, cfg *config.Config) {

	router.GET("/health", handler.Health(cfg))

}