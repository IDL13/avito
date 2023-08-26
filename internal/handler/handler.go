package handler

import (
	"net/http"

	request "github.com/IDL13/avito/internal/requests"
)

type Handler struct{}

func (h *Handler) StartServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server start"))
	// err := request.CreateTables()
	// if err != nil {
	// 	panic(err)
	// }
}

func (h *Handler) CreateSegment(w http.ResponseWriter, r *http.Request) {
	err := request.InserSegment("AVITO_DISCOUNT_30")
	if err != nil {
		w.Write([]byte("This segment is using"))
	}
}

func (h *Handler) DeletingSegment(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) AddingUserToSegment(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) GettingActiveUserSegments(w http.ResponseWriter, r *http.Request) {
}
