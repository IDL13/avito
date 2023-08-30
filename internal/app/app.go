package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/IDL13/avito/internal/CSV"
	"github.com/IDL13/avito/internal/handler"
)

type App struct {
	s *http.Server
	h handler.Handler
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
	a.h = handler.New()
	http.HandleFunc("/", a.h.StartServer)
	http.HandleFunc("/create_segment", a.h.CreateSegment)
	http.HandleFunc("/deleting_segment", a.h.DeletingSegment)
	http.HandleFunc("/adding_user_to_segment", a.h.AddDelSegments)
	http.HandleFunc("/getting_active_user_segments", a.h.GettingActiveUserSegments)
	http.HandleFunc("/ttl_adding_user_to_segment", a.h.TtlAddDelSegments)
	http.HandleFunc("/history", a.h.Hishtory)
	return a
}

func (a *App) Run(stop chan bool) {
	fmt.Println(`
╔══╗╔══╗─╔╗───╔╗╔══╗
╚╗╔╝║╔╗╚╗║║──╔╝║╚═╗║
─║║─║║╚╗║║║──╚╗║╔═╝║
─║║─║║─║║║║───║║╚═╗║
╔╝╚╗║╚═╝║║╚═╗─║║╔═╝║
╚══╝╚═══╝╚══╝─╚╝╚══╝
	`)
	fmt.Println("[SERVER STARTED]")
	fmt.Println("http://127.0.0.1:8080/")
	if err := CSV.CreateCSV(); err != nil {
		panic(err)
	}
	// if err := request.CreateTables(); err != nil {
	// 	panic(err)
	// }
	if err := a.s.ListenAndServe(); err != nil {
		stop <- true
		os.Exit(1)
	}
}
