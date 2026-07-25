package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/ronitmalvi/uber-backend/internal/config"
	"github.com/ronitmalvi/uber-backend/internal/handler"
)

func Register(router *gin.Engine, cfg *config.Config) {

	router.GET("/health", handler.Health(cfg))

}