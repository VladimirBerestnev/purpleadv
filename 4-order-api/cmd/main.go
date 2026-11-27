package main

import (
	"fmt"
	"net/http"
	"order-api/auth"
	"order-api/configs"
	"order-api/middleware"
	"order-api/pkg/db"
	"order-api/pkg/product"
	"order-api/pkg/user"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	repo := product.NewRepo(db)
	userRepo := *user.NewUserRepository(db)
	authService := auth.NewAuthService(&userRepo)
	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logger(router),
	}
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	product.NewProductHandler(router, product.ProductHandlerCfg{
		Repo:   repo,
		Config: conf,
	})
	fmt.Println("Server is running...")
	server.ListenAndServe()

}
