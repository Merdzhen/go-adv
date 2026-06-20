package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/hello"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf)

	router := http.NewServeMux()
	hello.NewHelloHandler(router)
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
