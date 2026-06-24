package auth

import (
	"go/adv-demo/configs"
	"go/adv-demo/pkg/request"
	"go/adv-demo/pkg/response"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
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
		_, err := request.HandleBody[RegisterRequest](w, req)
		if err != nil {
			return
		}

		data := RegisterResponse{
			Token: "123",
		}
		response.Json(w, data, 200)
	}
}
