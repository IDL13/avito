package main

import (
	"github.com/IDL13/avito/internal/app"
)

func main() {
	// sigs := make(chan os.Signal, 1)
	// done := make(chan bool, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// go func() {
	// 	done <- true
	// 	os.Remove("./csv_log.csv")
	// }()
	a := app.New()
	a.Run()
}
