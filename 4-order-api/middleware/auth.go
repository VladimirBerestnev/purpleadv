package middleware

import (
	"context"
	"errors"
	"net/http"
	"order-api/configs"
	"order-api/pkg/jwt"
	"strings"
)

func Auth(next http.HandlerFunc, conf *configs.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := bearerToken(r)
		if err != nil {
			http.Error(w, "нет токена", http.StatusUnauthorized)
			return
		}
		j := jwt.NewJWT(conf.Auth.Secret)
		phone, err := j.Parse(tokenString)
		if err != nil {
			http.Error(w, "невалидный токен", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "phone", phone)
		next(w, r.WithContext(ctx))
	}
}

func bearerToken(r *http.Request) (string, error) {
	const prefix = "Bearer "
	h := r.Header.Get("Authorization")
	if !strings.HasPrefix(h, prefix) {
		return "", errors.New("нет заголовка Authorization: Bearer")
	}
	return strings.TrimPrefix(h, prefix), nil
}
