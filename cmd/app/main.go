package main

import (
	"log"

	"github.com/elibr-edu/gateway/internal/app"
)

func main() {
	app := app.NewApp()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
	log.Println("Server exited")
}
