package link

import (
	"fmt"
	"go/adv-demo/pkg/request"
	"go/adv-demo/pkg/response"
	"net/http"
)

type LinkHandlerDeps struct{
	LinkRepository *LinkRepository
}

type LinkHandler struct{
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.Goto())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](w, req)
		if err != nil {
			return
		}

		// эти 2 сстроки должны быть в сервисном слое, но оставляем тут для простоты
		link := NewLink(body.Url) // теоретически может быть дублировние хэша при создании ссылки и это приведет к ошибке, пока это опускаем
		createdLink, err := handler.LinkRepository.Create(link)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdLink, 201)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")
		fmt.Printf("Update id: %s\n", id)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")
		fmt.Printf("Delete id: %s\n", id)
	}
}

func (handler *LinkHandler) Goto() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		fmt.Printf("Go to alias: %s\n", hash)
	}
}
