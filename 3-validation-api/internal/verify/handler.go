package verify

import (
	"fmt"
	"go/validation-api/config"
	"net/http"
)

type VerifyHandlerDeps struct {
	UserConfig *config.UserConfig
}

type VerifyHandler struct {
	userConfig *config.UserConfig
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		userConfig: deps.UserConfig,
	}
	router.HandleFunc("POST /send", handler.Send)
	router.HandleFunc("GET /verify/{hash}", handler.Verify)
}

// POST /send
func (h *VerifyHandler) Send(w http.ResponseWriter, req *http.Request) {
	fmt.Println("send")
}

// GET /verify/{hash}
func (h *VerifyHandler) Verify(w http.ResponseWriter, req *http.Request) {
	hash := req.PathValue("hash")
	fmt.Printf("verify triggered для хэша: %s\n", hash)
}
