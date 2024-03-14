package handler

import (
	"encoding/json"
	"log"
	"strings"

	"net/http"

	"github.com/anirudh97/GollabEdit/internal/service"
	utils "github.com/anirudh97/GollabEdit/pkg"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	log.Println("Handler | AuthMiddleware | Info :: Invoked")
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := utils.ValidateToken(tokenString)
		if err != nil {
			log.Println("Handler | AuthMiddleware | Error :: Invalid Token")
			resp := &utils.Response{
				Data:  nil,
				Error: service.ErrInvalidToken.Error(),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		// Token is valid, set the user information from token into the context
		if claims, ok := token.Claims.(*utils.Claims); ok && token.Valid {
			c.Set("email", claims.Email)
			c.Next()
		} else {
			log.Println("Handler | AuthMiddleware | Error :: Invalid Token")
			resp := &utils.Response{
				Data:  nil,
				Error: service.ErrInvalidToken.Error(),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		}
	}
}

// Parses the request data and calls the CreateUser Service.
func CreateUser(c *gin.Context) {
	log.Println("Handler | CreateUser | Info :: Invoked")

	var req service.CreateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Println("Handler | CreateUser | Error :: Error: ", err.Error())
		resp := &utils.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// Email validation
	if !utils.ValidateEmail(req.Email) {
		log.Println("Handler | CreateUser | Info :: Email Validation Failed")
		resp := &utils.Response{
			Data:  nil,
			Error: "Email not in the correct format.",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Password validation
	if !utils.ValidatePassword(req.Password) {
		log.Println("Handler | CreateUser | Info :: Password Validation Failed")
		resp := &utils.Response{
			Data:  nil,
			Error: "Password does not meet the requirements. It should be atleast 8 characters long and alpha numeric.",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user, err := service.CreateUser(&req)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			log.Println("Handler | CreateUser | Info :: User already exists")
			resp := &utils.Response{
				Data:  nil,
				Error: service.ErrUserAlreadyExists.Error(),
			}
			c.JSON(http.StatusConflict, resp)
		} else {
			log.Println("Handler | CreateUser | Error :: Error: ", err.Error())
			resp := &utils.Response{
				Data:  nil,
				Error: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, resp)
		}
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Println("Handler | CreateUser | Error :: Error: ", err.Error())
		resp := &utils.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := &utils.Response{
		Data:  jsonData,
		Error: "",
	}
	c.JSON(http.StatusCreated, resp)

}

func Login(c *gin.Context) {
	log.Println("Handler: Login: Invoked")

	var req service.LoginUserRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Println("Handler | Login | Error :: Error: ", err.Error())
		resp := &utils.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// Email validation
	log.Println(req.Email)
	if !utils.ValidateEmail(req.Email) {
		log.Println("Handler | Login | Info :: Email Validation Failed")
		resp := &utils.Response{
			Data:  nil,
			Error: "Email not in the correct format.",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	loggedInUser, loginErr := service.LoginUser(&req)
	if loginErr != nil {
		if loginErr == service.ErrUserDoesNotExist {
			log.Println("Handler | Login | Info :: User does not exist")
			resp := &utils.Response{
				Data:  nil,
				Error: service.ErrUserDoesNotExist.Error(),
			}
			c.JSON(http.StatusConflict, resp)
		} else if loginErr == service.ErrIncorrectPassword {
			log.Println("Handler | Login | Info :: Incorrect Password")
			resp := &utils.Response{
				Data:  nil,
				Error: service.ErrIncorrectPassword.Error(),
			}
			c.JSON(http.StatusConflict, resp)
		} else {
			log.Println("Handler | Login | Error :: Error: ", loginErr.Error())
			resp := &utils.Response{
				Data:  nil,
				Error: loginErr.Error(),
			}
			c.JSON(http.StatusInternalServerError, resp)
		}
		return
	}

	jsonData, err := json.Marshal(loggedInUser)
	if err != nil {
		log.Println("Handler | Login | Error :: Error: ", err.Error())
		resp := &utils.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := &utils.Response{
		Data:  jsonData,
		Error: "",
	}

	c.JSON(http.StatusOK, resp)

}

func Logout(c *gin.Context) {
	log.Println("Handler: Logout: Invoked")

	c.JSON(http.StatusOK, nil)
}
