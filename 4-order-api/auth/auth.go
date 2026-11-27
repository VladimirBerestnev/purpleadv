package auth

import (
	"net/http"
	"order-api/configs"
	"order-api/pkg/user"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
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
