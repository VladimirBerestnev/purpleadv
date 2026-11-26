package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/random", homeRandom)

	http.ListenAndServe(":8080", router)
}

func homeRandom(w http.ResponseWriter, req *http.Request) {
	number := rand.IntN(6) + 1
	numberByte, err := json.Marshal(number)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(numberByte)
}
