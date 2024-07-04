package model

import (
	"fmt"
	"os"
	"sosservice/src/configurations/rest_err"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	JWT_KEY = "JWT_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_KEY)

	// TODO Encrypt data before put it on claims

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("Error generating JWT token: %s", err.Error()))
	}

	return tokenString, nil
}
