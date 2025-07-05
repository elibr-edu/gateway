package main

import (
	"fmt"
	"log"

	"github.com/elibr-edu/gateway/internal/app"
	"github.com/elibr-edu/gateway/pkg/config"
)

func main() {
	cfg := config.MustLoadWithEnv()

	app := app.NewApp(cfg)

	fmt.Printf(`
=============================================

Server is running on %s

=============================================
Configuration:

%s

=============================================
`, cfg.Server.ServerAddr(), cfg.Format())

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
	log.Println("Server exited")
}
