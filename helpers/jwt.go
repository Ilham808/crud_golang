package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(secret string, userId uint) (string, error) {
	secretKey := []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secretKey)
}

func GenerateRefreshToken(refreshSecret string, userId uint) (string, error) {
	refreshSecretKey := []byte(refreshSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	return token.SignedString(refreshSecretKey)
}

func VerifyToken(secret string, tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}

func VerifyRefreshToken(refreshSecret string, tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})
}
