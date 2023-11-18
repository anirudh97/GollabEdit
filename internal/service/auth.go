package service

import (
	"github.com/anirudh97/GollabEdit/internal/model"
	"github.com/anirudh97/GollabEdit/internal/repository"
	utils "github.com/anirudh97/GollabEdit/pkg"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type CreateUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginUserRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type LoginUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func CreateUser(r *CreateUserRequest) (*CreateUserResponse, error) {

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return nil, hashErr
	}

	user := &model.User{
		Username: r.Username,
		Email:    r.Email,
		Password: string(hashedPassword),
	}

	// Check if user already exists
	status, err := repository.CheckForUserExistence(user.Email)
	if err != nil {
		return nil, err
	}
	if status {
		return nil, ErrUserAlreadyExists
	}

	// create user
	createUserErr := repository.CreateUser(user)
	if createUserErr != nil {
		return nil, createUserErr
	}

	cu := &CreateUserResponse{
		Username: r.Username,
		Email:    r.Email,
	}
	return cu, nil
}

func verifyPassword(ip string, ep string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(ep), []byte(ip))
	return err == nil

}

func LoginUser(r *LoginUserRequest) (*model.LoggedInUser, error) {
	// Check if user already exists

	status, err := repository.CheckForUserExistence(r.Email)
	if err != nil {
		return nil, err
	}
	if !status {
		return nil, ErrUserDoesNotExist
	}

	dbOutput, dbFetchErr := repository.LoginUser(r.Email)
	if dbFetchErr != nil {
		return nil, dbFetchErr
	}

	if !verifyPassword(r.Password, dbOutput.Password) {
		return nil, ErrIncorrectPassword
	}

	// Generate JWT token
	jwtToken, tokenErr := utils.GenerateJWTToken(dbOutput.Email)
	if tokenErr != nil {
		return nil, tokenErr
	}

	loggedInUser := &model.LoggedInUser{
		Username: dbOutput.Username,
		Email:    dbOutput.Email,
		Token:    jwtToken,
	}
	return loggedInUser, nil
}
