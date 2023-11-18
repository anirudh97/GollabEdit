package repository

import (
	"github.com/anirudh97/GollabEdit/internal/database"
	"github.com/anirudh97/GollabEdit/internal/model"
)

func CheckForUserExistence(e string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"

	var exists bool

	err := database.DB.Get(&exists, query, e)

	return exists, err

}

func CreateUser(u *model.User) error {
	query := "INSERT INTO users (username, email, hashPassword) VALUES (?, ?, ?)"

	_, err := database.DB.Exec(query, u.Username, u.Email, u.Password)

	return err
}

func LoginUser(email string) (*model.User, error) {
	query := "SELECT * FROM users WHERE email=?"
	var user model.User
	err := database.DB.Get(&user, query, email)

	return &user, err
}
