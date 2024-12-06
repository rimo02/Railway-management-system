package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var secret_key = "my_secret_key"

func GenerateJWT(userID uint, role string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userID": userID,
        "role":   role,
        "exp":    time.Now().Add(time.Hour * 48).Unix(),
    })
    return token.SignedString([]byte(secret_key))
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret_key), nil
    })
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, err
}

func GetUserIdFromJWTToken(c *gin.Context) (uint, error) {

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return 0, errors.New("missing Authorization Header")
	}

	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return 0, errors.New("invalid token")
	}
	userID, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("invalid userID in token")
	}

	return uint(userID), nil
}
