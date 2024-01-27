package service

import "errors"

// Predefined errors for the service layer.
var (
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrUserDoesNotExist    = errors.New("user does not exist. create an account")
	ErrIncorrectPassword   = errors.New("incorrect password")
	ErrInvalidToken        = errors.New("invalid token. login again")
	ErrFileFormatIncorrect = errors.New("invalid file format or no format mentioned. only txt and text format allowed")
	ErrFileAlreadyExists   = errors.New("file already exists")
	ErrFileDoesNotExist    = errors.New("file does not exist. please create one")
	ErrAlreadyShared       = errors.New("file already shared")
)
