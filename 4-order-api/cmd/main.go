package main

import (
	"fmt"
	"net/http"
	"order-api/configs"
	"order-api/middleware"
	"order-api/pkg/db"
	"order-api/pkg/product"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	conf := configs.LoadConfig()
	db := db.NewDb(conf)

	repo := product.NewRepo(db)

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logger(router),
	}

	product.NewProductHandler(router, product.ProductHandlerCfg{
		Repo: repo,
	})
	fmt.Println("Server is running...")
	server.ListenAndServe()

}
