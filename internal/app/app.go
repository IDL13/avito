package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IDL13/avito/internal/handler"
)

type App struct {
	s *http.Server
	h *handler.Handler
}

func New() *App {
	a := &App{
		s: &http.Server{
			Addr:           ":8080",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
	http.HandleFunc("/", a.h.StartServer)
	http.HandleFunc("/create_segment", a.h.CreateSegment)
	http.HandleFunc("/deleting_segment", a.h.DeletingSegment)
	http.HandleFunc("/adding_user_to_segment", a.h.AddingUserToSegment)
	http.HandleFunc("/getting_active_user_segments", a.h.GettingActiveUserSegments)
	return a
}

func (a *App) Run() {
	fmt.Println("[SERVER STARTED]")
	log.Fatal(a.s.ListenAndServe())
}
