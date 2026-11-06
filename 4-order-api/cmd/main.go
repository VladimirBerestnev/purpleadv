package main

import (
	"fmt"
	"net/http"
	"order-api/configs"
	"order-api/pkg/db"
	"order-api/verify"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	verify.VerifyHandler(router)
	fmt.Println("Server is running...")
	server.ListenAndServe()

}
