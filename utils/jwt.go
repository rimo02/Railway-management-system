package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var access_secret_key = "my_access_secret_key"
var refresh_secret_key = "my_refresh_secret_key"

func GenerateAccessToken(userUUID string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userUUID": userUUID,
		"role":     role,
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // access token valid for 15 mins
	})
	return token.SignedString([]byte(access_secret_key))
}

func GenerateRefreshToken(userUUID string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userUUID": userUUID,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // refresh token valid for 1 wweek
	})
	return token.SignedString([]byte(refresh_secret_key))
}

func ValidateJWT(tokenString string, isRefresh bool) (jwt.MapClaims, error) {
	secretKey := access_secret_key
	if isRefresh {
		secretKey = refresh_secret_key
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func GetUserIdFromJWTToken(c *gin.Context) (string, error) {

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return "", errors.New("missing Authorization Header")
	}

	claims, err := ValidateJWT(tokenString, false) // This is for access token
	if err != nil {
		return "", errors.New("invalid token")
	}
	userUUID, ok := claims["userUUID"].(string)
	if !ok {
		return "", errors.New("invalid userID in token")
	}

	return userUUID, nil
}

// Hash password before storing in DB
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Verify hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
