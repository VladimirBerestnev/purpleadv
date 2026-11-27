package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func (j *JWT) Parse(tokenString string) (string, error) {
	token, err := jwt.Parse(
		tokenString,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("неожиданный алгоритм подписи: %v", t.Header["alg"])
			}
			return []byte(j.Secret), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("токен невалиден")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("не удалось прочитать claims")
	}

	phone, ok := claims["phone"].(string)
	if !ok || phone == "" {
		return "", errors.New("в токене нет phone")
	}

	return phone, nil
}
