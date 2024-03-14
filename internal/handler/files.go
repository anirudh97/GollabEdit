package handler

import (
	"encoding/json"
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
		log.Println("Handler | CreateFile | Error :: Error: ", err.Error())
		resp := &utils.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// Filename validation
	if !utils.ValidateFilename(req.Filename) {
		log.Println("Handler | CreateFile | Info :: Filename validation failed ")
		resp := &utils.Response{
			Data:  nil,
			Error: service.ErrFileFormatIncorrect.Error(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	respCreateFile, err := service.CreateFile(&req)
	if err != nil {
		if err == service.ErrFileAlreadyExists {
			log.Println("Handler | CreateFile | Info :: File already exists ")
			resp := &utils.Response{
				Data:  nil,
				Error: service.ErrFileAlreadyExists.Error(),
			}
			c.JSON(http.StatusConflict, resp)
		} else {
			log.Println("Handler | CreateFile | Error :: Error: ", err.Error())
			resp := &utils.Response{
				Data:  nil,
				Error: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, resp)
		}
		return
	}

	jsonData, err := json.Marshal(respCreateFile)
	if err != nil {
		log.Println("Handler | CreateFile | Error :: Error: ", err.Error())
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

func OpenFile(c *gin.Context) {
	log.Println("Handler | OpenFile :: Invoked")
	var req service.OpenFileRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("Handler | OpenFile | Error :: Error: ", err.Error())
		resp := &utils.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	respOpen, err := service.OpenFile(&req)
	if err != nil {
		if err == service.ErrFileDoesNotExist {
			log.Println("Handler | OpenFile | Info :: File Does not exist ")
			resp := &utils.Response{
				Data:  nil,
				Error: service.ErrFileDoesNotExist.Error(),
			}
			c.JSON(http.StatusConflict, resp)

		} else {
			log.Println("Handler | OpenFile | Error :: Error: ", err.Error())
			resp := &utils.Response{
				Data:  nil,
				Error: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, resp)
		}
		return
	}
	jsonData, err := json.Marshal(respOpen)
	if err != nil {
		log.Println("Handler | OpenFile | Error :: Error: ", err.Error())
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

func ShareFile(c *gin.Context) {
	log.Println("Handler | ShareFile :: Invoked")

	var req service.ShareFileRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println("Handler | ShareFile | Error :: Error: ", err.Error())
		resp := &utils.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// Email validation
	if !utils.ValidateEmail(req.SharedByEmail) {
		log.Println("Handler | ShareFile | Info :: Email Validation Failed")
		resp := &utils.Response{
			Data:  nil,
			Error: "Email not in the correct format.",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if !utils.ValidateEmail(req.SharedWithEmail) {
		log.Println("Handler | ShareFile | Info :: Email Validation Failed")
		resp := &utils.Response{
			Data:  nil,
			Error: "Email not in the correct format.",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	respShare, err := service.ShareFile(&req)
	if err != nil {
		if err == service.ErrUserDoesNotExist {
			log.Println("Handler | ShareFile | Info :: Invalid User")
			resp := &utils.Response{
				Data:  nil,
				Error: service.ErrUserDoesNotExist.Error(),
			}
			c.JSON(http.StatusBadRequest, resp)
		} else if err == service.ErrAlreadyShared {
			log.Println("Handler | ShareFile | Info :: Already Shared")
			resp := &utils.Response{
				Data:  nil,
				Error: service.ErrAlreadyShared.Error(),
			}
			c.JSON(http.StatusConflict, resp)
		} else {
			log.Println("Handler | ShareFile | Error :: Error: ", err.Error())
			resp := &utils.Response{
				Data:  nil,
				Error: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, resp)

		}
		return
	}

	jsonData, err := json.Marshal(respShare)
	if err != nil {
		log.Println("Handler | ShareFile | Error :: Error: ", err.Error())
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

func InsertCharacter(c *gin.Context) {
	log.Println("Handler | InsertCharacter :: Invoked")
	var req service.InsertCharacterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Handler | InsertCharacter | Error :: Error: ", err.Error())
		resp := &utils.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	err := service.InsertCharacter(&req)
	if err != nil {
		log.Println("Handler | InsertCharacter | Error :: Error: ", err.Error())
		resp := &utils.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := &utils.Response{
		Data:  nil,
		Error: "",
	}
	c.JSON(http.StatusCreated, resp)

}

func DeleteCharacter(c *gin.Context) {
	log.Println("Handler | DeleteCharacter :: Invoked")
	var req service.DeleteCharacterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Handler | DeleteCharacter | Error :: Error: ", err.Error())
		resp := &utils.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	err := service.DeleteCharacter(&req)
	if err != nil {
		log.Println("Handler | DeleteCharacter | Error :: Error: ", err.Error())
		resp := &utils.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := &utils.Response{
		Data:  nil,
		Error: "",
	}
	c.JSON(http.StatusCreated, resp)
}
