package middleware

import (
	"context"
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)


func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		fmt.Println("isValid", isValid)
		fmt.Println("data", data)

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
