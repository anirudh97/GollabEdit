package handler

import (
	"log"

	"net/http"

	"github.com/anirudh97/GollabEdit/internal/service"
	utils "github.com/anirudh97/GollabEdit/pkg"
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

	// Email validation
	if !utils.ValidateEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is not in the right format"})
		return
	}

	// Password validation
	if !utils.ValidatePassword(req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password does not meet the requirements"})
		return
	}

	user, err := service.CreateUser(&req)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": service.ErrUserAlreadyExists.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, user)

}

func Login(c *gin.Context) {
	log.Println("Handler: CreateUser: Invoked")

	var req service.LoginUserRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Println("Error in Binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Email validation
	if !utils.ValidateEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is not in the right format"})
		return
	}

	loggedInUser, loginErr := service.LoginUser(&req)
	if loginErr != nil {
		if loginErr == service.ErrUserDoesNotExist {
			c.JSON(http.StatusConflict, gin.H{"error": service.ErrUserDoesNotExist.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": loginErr.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, loggedInUser)

}

func Logout(c *gin.Context) {
	log.Println("Handler: CreateUser: Invoked")
}
