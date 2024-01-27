package api

import (
	"github.com/anirudh97/GollabEdit/internal/handler"
	"github.com/gin-gonic/gin"
)

func fileRoutes(router *gin.Engine) {
	router.POST("/file", handler.AuthMiddleware(), handler.CreateFile)
	router.POST("/file/open", handler.AuthMiddleware(), handler.OpenFile)
	// router.DELETE("/file", handler.DeleteFile)
	// router.POST("/file", handler.GetFile)

}
