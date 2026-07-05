package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	database := db.NewDb(conf)
	if err := database.CheckConnection(); err != nil {
		fmt.Printf("БД недоступна: %v\n", err)
	} else {
		fmt.Println("Успешное подключение к базе данных!")
	}

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
