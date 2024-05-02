package api

import (
	"github.com/anirudh97/GollabEdit/internal/handler"
	"github.com/gin-gonic/gin"
)

func fileRoutes(router *gin.Engine) {
	router.POST("/file", handler.AuthMiddleware(), handler.CreateFile)
	router.POST("/file/open", handler.AuthMiddleware(), handler.OpenFile)
	router.POST("/file/share", handler.AuthMiddleware(), handler.ShareFile)
	router.POST("/ws", handler.AuthMiddleware(), handler.WebsocketHandler)

	// router.DELETE("/file", handler.DeleteFile)
	// router.POST("/file", handler.GetFile)

	// CRDT-related endpoints
	// router.POST("/file/insert", handler.AuthMiddleware(), handler.InsertCharacter)
	// router.DELETE("/file/delete", handler.AuthMiddleware(), handler.DeleteCharacter)

}
