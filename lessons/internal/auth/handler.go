package auth

import (
	"go/adv-demo/configs"
	"go/adv-demo/pkg/request"
	"go/adv-demo/pkg/response"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("/auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		_, err := request.HandleBody[LoginRequest](w, req)
		if err != nil {
			return
		}

		data := LoginResponse{
			Token: "123",
		}
		response.Json(w, data, 201)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](w, req)
		if err != nil {
			return
		}
		_, err = handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			if err.Error() == ErrUserExists {
				http.Error(w, "User already exists", http.StatusBadRequest)
				return
			}
			http.Error(w, "Internal server error", http.StatusBadRequest)
			return
		}
	}
}
