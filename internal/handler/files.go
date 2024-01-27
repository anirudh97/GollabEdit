package handler

import (
	"log"

	"net/http"

	"github.com/anirudh97/GollabEdit/internal/service"
	utils "github.com/anirudh97/GollabEdit/pkg"
	"github.com/gin-gonic/gin"
)

func CreateFile(c *gin.Context) {
	log.Println("Handler | CreateFile :: Invoked")

	var req service.CreateFileRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("Error in Binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Filename validation
	if !utils.ValidateFilename(req.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": service.ErrFileFormatIncorrect.Error()})
		return
	}

	resp, err := service.CreateFile(&req)
	if err != nil {
		if err == service.ErrFileAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": service.ErrFileAlreadyExists.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func OpenFile(c *gin.Context) {
	log.Println("Handler | OpenFile :: Invoked")
	var req service.OpenFileRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("Error in Binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := service.OpenFile(&req)
	if err != nil {
		if err == service.ErrFileDoesNotExist {
			c.JSON(http.StatusConflict, gin.H{"error": service.ErrFileDoesNotExist.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, resp)

}

func ShareFile(c *gin.Context) {
	log.Println("Handler | ShareFile :: Invoked")

	var req service.ShareFileRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("Error in Binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Email validation
	if !utils.ValidateEmail(req.SharedByEmail) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is not in the right format"})
		return
	}

	if !utils.ValidateEmail(req.SharedWithEmail) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is not in the right format"})
		return
	}

	resp, err := service.ShareFile(&req)
	if err != nil {
		if err == service.ErrUserDoesNotExist {
			c.JSON(http.StatusConflict, gin.H{"error": service.ErrUserDoesNotExist.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, resp)
}
