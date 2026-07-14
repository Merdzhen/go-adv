package main

import (
	"context"
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
	"go/adv-demo/pkg/middleware"
	"net/http"
	"time"
)

func tickOperation(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <- ticker.C:
			fmt.Println("tick")
		case <- ctx.Done():
			fmt.Println("cancel")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go tickOperation(ctx)

	time.Sleep(2 * time.Second)
	cancel()
}

func main2() {
	conf := configs.LoadConfig()

	database := db.NewDb(conf)
	if err := database.CheckConnection(); err != nil {
		fmt.Printf("БД недоступна: %v\n", err)
	} else {
		fmt.Println("Успешное подключение к базе данных!")
	}
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)

	// Services
	authservice := auth.NewAuthService(userRepository)

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authservice,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
