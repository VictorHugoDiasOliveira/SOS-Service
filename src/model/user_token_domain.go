package model

import (
	"fmt"
	"os"
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func ValidateTokenMiddleware(context *gin.Context) {
	secret := os.Getenv(JWT_KEY)

	tokenValue := RemoveBearerPrefix(context.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		context.Abort()
		return nil, rest_err.NewBadRequestError("Invalid Token")
	})
	if err != nil {
		errRest := rest_err.NewUnauthorizedRequestError("Invalid Token")
		context.JSON(errRest.Code, errRest)
		context.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedRequestError("Invalid Token")
		context.JSON(errRest.Code, errRest)
		context.Abort()
		return
	}

	userDomain := &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int(claims["age"].(float64)),
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", userDomain))
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix("Bearer ", token)
	}
	return token
}
