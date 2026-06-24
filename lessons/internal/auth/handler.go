package auth

import (
	"encoding/json"
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/res"
	"net/http"
	"net/mail"
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
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}

		if payload.Email == "" {
			res.Json(w, "Email required", 400)
			return
		}

		// способ 1
		// reg, _ := regexp.Compile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`)
		// if !reg.MatchString(payload.Email) {
		// 	res.Json(w, "Wrong email", 400)
		//   return
		// }

		// способ 2
		// match, _ := regexp.MatchString(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`, payload.Email)
		// if !match {
		// 	res.Json(w, "Wrong email", 400)
		//   return
		// }

		// способ 3
		mailAddress, err := mail.ParseAddress(payload.Email)
		if err != nil {
			res.Json(w, "Wrong email", 400)
			return
		}
		fmt.Println(mailAddress)

		if payload.Password == "" {
			res.Json(w, "Password required", 400)
			return
		}

		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, 201)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("register")
	}
}
