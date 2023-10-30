package handler

import (
	"log"

	"net/http"

	"github.com/anirudh97/GollabEdit/internal/service"
	"github.com/gin-gonic/gin"
)

// Parses the request data and calls the CreateUser Service.
func CreateUser(c *gin.Context) {
	log.Println("Handler | CreateUser :: Invoked")

	var req service.CreateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Println("Error in Binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(req)
	user, err := service.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new user"})
		return
	}

	c.JSON(http.StatusCreated, user)

}

func Login(c *gin.Context) {
	log.Println("Handler: CreateUser: Invoked")
}

func Logout(c *gin.Context) {
	log.Println("Handler: CreateUser: Invoked")
}
