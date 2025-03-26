package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId uint `json:"user_id"`
}

func GenerateJwtToken(userId uint) (string, error) {
	payload := &TokenClaims{}
	payload.ExpiresAt = &jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(3))}
	payload.IssuedAt = &jwt.NumericDate{Time: time.Now()}
	payload.UserId = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
