package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/stat"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
	"go/adv-demo/pkg/event"
	"go/adv-demo/pkg/middleware"
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
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)
	statRepository := stat.NewStatRepository(database)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(stat.StatServiceDeps{
		EventBus: eventBus,
		StatRepository: statRepository,
	})

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus: eventBus,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		Config: conf,
		StatRepository: statRepository,
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

	go statService.AddClick()

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
