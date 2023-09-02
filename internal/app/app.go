package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/IDL13/avito/docs"
	"github.com/IDL13/avito/internal/CSV"
	"github.com/IDL13/avito/internal/handler"
	httpSwager "github.com/swaggo/http-swagger"
)

type App struct {
	s   *http.Server
	h   handler.Handler
	mux *http.ServeMux
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
	a.mux = http.NewServeMux()
	a.s.Handler = a.mux
	a.mux.HandleFunc("/", a.h.StartServer)
	a.mux.HandleFunc("/create_segment", a.h.CreateSegment)
	a.mux.HandleFunc("/deleting_segment", a.h.DeletingSegment)
	a.mux.HandleFunc("/create_user", a.h.CreateUser)
	a.mux.HandleFunc("/deleting_user", a.h.DeletingUser)
	a.mux.HandleFunc("/adding_user_to_segment", a.h.AddDelSegments)
	a.mux.HandleFunc("/getting_active_user_segments", a.h.GettingActiveUserSegments)
	a.mux.HandleFunc("/ttl_adding_user_to_segment", a.h.TtlAddDelSegments)
	a.mux.HandleFunc("/history", a.h.Hishtory)
	a.mux.HandleFunc("/docs/", httpSwager.WrapHandler)
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
	fmt.Println("http://127.0.0.1:8080/docs/")
	if err := CSV.CreateCSV(); err != nil {
		fmt.Fprintf(os.Stderr, "CSV file not created:%v", err)
		os.Exit(1)
	}
	// if err := request.CreateTables(); err != nil {
	// 	panic(err)
	// }
	if err := a.s.ListenAndServe(); err != nil {
		stop <- true
		os.Exit(1)
	}
}
