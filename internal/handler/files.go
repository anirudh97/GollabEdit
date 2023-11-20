package handler

import (
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFile(c *gin.Context) {
	log.Println("Handler | CreateFile :: Invoked")
	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}
