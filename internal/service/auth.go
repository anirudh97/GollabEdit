package service

import (
	"github.com/anirudh97/GollabEdit/internal/model"
)

type CreateUserRequest struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func CreateUser(r *CreateUserRequest) (*model.User, error) {
	user := &model.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}

	return user, nil
	// return repository.CreateUser(user)
}
