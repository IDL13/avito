package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IDL13/avito/internal/app"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	stop := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		os.Remove("./csv_log.csv")
		done <- true
	}()
	a := app.New()
	go a.Run(stop)
	select {
	case stop_r := <-done:
		fmt.Println(stop_r)
		os.Exit(1)
	case stop_l := <-stop:
		fmt.Println(stop_l)
		os.Exit(1)
	}
}
