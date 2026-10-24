package main

import (
	"encoding/json"
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
	number := rand.IntN(6) + 1
	numberByte, err := json.Marshal(number)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(numberByte)
}
