package utils

import (
	"regexp"
	"strings"
	"time"

	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	jwtKey          = []byte("GollabEdit")
	ErrInvalidToken = errors.New("invalid token. login again")
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func ValidateFilename(filename string) bool {
	split := strings.Split(filename, ".")

	if len(split) != 2 {
		return false
	} else if split[1] != "txt" && split[1] != "text" {
		return false
	}

	return true

}

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(email)
}

func ValidatePassword(password string) bool {
	passwordRegex := regexp.MustCompile(`^[A-Za-z0-9]{8,}$`)

	return passwordRegex.MatchString(password)
}

func GenerateJWTToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return token, nil
}
