package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(email string) (string, error) {
	segredo := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 1).Unix(),
		})
	tokenString, err := token.SignedString([]byte(segredo))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func VerifyJwtToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}
	return nil
}
