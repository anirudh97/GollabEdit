package model

type User struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"hashPassword"`
}

type LoggedInUser struct {
	Username string
	Email    string
	Token    string
}
