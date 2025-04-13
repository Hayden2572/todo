package main

import (
	todo "ToDo"
	"ToDo/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while rinning http server: %s", err.Error())
	}
}
