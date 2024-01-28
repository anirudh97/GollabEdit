package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

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

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
	PATCH  HTTPMethod = "PATCH"
)

type RequestConfig struct {
	BaseURL string
	API     string
	Method  HTTPMethod
	Payload []byte
	Bearer  string
}

type Response struct {
	Data  []byte `json:"data"`
	Error string `json:"error"`
}

func DefaultRequestConfig() RequestConfig {
	return RequestConfig{
		BaseURL: "http://localhost:8080",
	}
}

func MakeRequest(r RequestConfig) (responseBody []byte, statusCode int, err error) {
	req, err := http.NewRequest(string(r.Method), r.BaseURL+r.API, bytes.NewBuffer(r.Payload))
	if err != nil {
		fmt.Println("Error creating request")
		return nil, 500, err
	}

	if len(r.Bearer) > 0 {
		req.Header.Add("Authorization", "Bearer "+r.Bearer)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error in response")
		return nil, 500, err
	}

	defer resp.Body.Close()
	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	return responseBody, resp.StatusCode, nil

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
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return "", err
	}

	expirationTime := time.Now().In(loc).Add(24 * time.Hour)

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
