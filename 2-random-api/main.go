package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
)

type RandomHandler struct {}

func NewRandomHandler(router *http.ServeMux) {
	handler := &RandomHandler{}
	router.HandleFunc("/random", handler.GetRandomNumber())
}

func (handler *RandomHandler) GetRandomNumber() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		randomNum := rand.IntN(6) + 1
		numStr := strconv.Itoa(randomNum)
		w.Write([]byte(numStr))
	}
}

func main() {
	router := http.NewServeMux()
	NewRandomHandler(router)

	server := http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8082")
	server.ListenAndServe()
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
