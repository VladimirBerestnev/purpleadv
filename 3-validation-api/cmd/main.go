package main

import (
	"net/http"
	"validation/verify"
)

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8082",
		Handler: router,
	}
	verify.VerifyHandler(router)

	server.ListenAndServe()
}
