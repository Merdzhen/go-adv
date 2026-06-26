package main

import (
	"fmt"
	"go/validation-api/config"
	"go/validation-api/internal/verify"
	"net/http"
)

func main() {
	conf := &config.Config{
		User: config.UserConfig{
			Address: "localhost:1025",
			Email:   "test@test.com",
			Password: "",
		},
	}

	router := http.NewServeMux()
	
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		UserConfig: &conf.User,
	})

	server := http.Server{
		Addr:    ":8083",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8083")
	server.ListenAndServe()
}
