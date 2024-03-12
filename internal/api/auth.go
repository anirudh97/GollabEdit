package api

import (
	"github.com/anirudh97/GollabEdit/internal/handler"
	"github.com/gin-gonic/gin"
)

func authRoutes(router *gin.Engine) {
	router.POST("/auth/create", handler.CreateUser)
	router.POST("/auth/login", handler.Login)
	router.POST("/auth/logout", handler.Logout)

}
