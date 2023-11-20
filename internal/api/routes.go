package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(g *gin.Engine) {
	// Auth routes
	authRoutes(g)
	//File routes
	fileRoutes(g)
}
