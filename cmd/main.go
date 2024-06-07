package main

import (
	"github.com/imjowend/microservicio-usuarios-y-autenticacion/cmd/api"
	"log"
)

func main() {

	log.Println("Starting application...")

	config, err := api.Config()
	if err != nil {
		log.Fatal("error at dependencies building", err)
	}

	app := api.Build(config)
	if err := app.Run(); err != nil {
		log.Fatal("error at app startup", err)
	}
}
