package service

import "errors"

// Predefined errors for the service layer.
var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserDoesNotExist  = errors.New("user does not exist. create an account")
	ErrIncorrectPassword = errors.New("incorrect password")
)
