package link

import (
	"fmt"
	"net/http"
)

type LinkHandlerDeps struct{}

type LinkHandler struct{}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{alias}", handler.Goto())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Create Link")
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		alias := req.PathValue("id")
		fmt.Printf("Update id: %s\n", alias)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		alias := req.PathValue("id")
		fmt.Printf("Delete id: %s\n", alias)
	}
}

func (handler *LinkHandler) Goto() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		alias := req.PathValue("alias")
		fmt.Printf("Go to alias: %s\n", alias)
	}
}
