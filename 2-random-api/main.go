package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	router.HandleFunc("/random", homeRandom)

	server.ListenAndServe()
}

func homeRandom(w http.ResponseWriter, req *http.Request) {
	fmt.Println(rand.IntN(7))
}
