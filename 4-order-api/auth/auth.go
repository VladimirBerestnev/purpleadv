package auth

import (
	"context"
	"errors"
	"net/http"
	"order-api/configs"
	"order-api/pkg/jwt"
	"order-api/pkg/user"
	"strings"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
	*jwt.JWT
}

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthService struct {
	userRepo *user.UserRepository
}

type AuthCode struct {
	SessionID string `json:"sessionId" validate:"required"`
	Code      int    `json:"code" validate:"required"`
}

type SessionIdResponse struct {
	SessionID string `json:"sessionId" validate:"required"`
}

func NewAuthService(userRepo *user.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) *AuthHandler {
	authHandler := AuthHandler{Config: deps.Config, AuthService: deps.AuthService}
	router.HandleFunc("POST /register", authHandler.GetCode())
	router.HandleFunc("POST /token", authHandler.GetToken())

	return &authHandler
}

func (authHandler *AuthHandler) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := bearerToken(r)
		if err != nil {
			http.Error(w, "нет токена", http.StatusUnauthorized)
			return
		}

		phone, err := authHandler.JWT.Parse(tokenString)
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
