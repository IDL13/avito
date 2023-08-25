package handler

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) StartServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server start"))
}

func (h *Handler) CreateSegment(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) DeletingSegment(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) AddingUserToSegment(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) GettingActiveUserSegments(w http.ResponseWriter, r *http.Request) {
}
