package stat

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/middleware"
	"net/http"
	"strings"
	"time"
)

const (
	FilterByDay   = "day"
	FilterByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}

	router.Handle("GET /stat", middleware.IsAuthed(handler.GetAll(), deps.Config))
}

func (handler *StatHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fromStr := req.URL.Query().Get("from")
		fromDate, err := time.Parse(time.DateOnly, fromStr)
		if err != nil && fromStr != "" {
			http.Error(w, "Invalid 'from' date format. Expected YYYY-MM-DD", http.StatusBadRequest)
			return
		}

		toStr := req.URL.Query().Get("to")
		toDate, err := time.Parse(time.DateOnly, toStr)
		if err != nil && fromStr != "" {
			http.Error(w, "Invalid 'to' date format. Expected YYYY-MM-DD", http.StatusBadRequest)
			return
		}

		by := strings.ToLower(req.URL.Query().Get("by"))
		if by != FilterByDay && by != FilterByMonth && fromStr != "" {
			http.Error(w, "Invalid 'by' param. Expected day/month", http.StatusBadRequest)
			return
		}

		fmt.Println("StatHandler", fromDate, toDate, by)
	}
}
