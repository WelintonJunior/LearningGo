package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "chaveultrasecreta3000"

func GenerateToken(cliEmail string, cliId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cliEmail": cliEmail,
		"cliId":    cliId,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
