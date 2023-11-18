package utils

import (
	"regexp"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("GollabEdit")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
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
