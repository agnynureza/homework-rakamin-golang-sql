package utils

import (
	"time"

	"github.com/agnynureza/homework-rakamin-golang-sql/config"
	"github.com/golang-jwt/jwt"
)

func GenerateNewAccessToken() (string, error) {
	secret := config.GetJWTSecret()
	minutesCount := config.GetJWTExp()

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
