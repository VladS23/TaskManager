package main

import (
	"log"
	todo "myTaskManager"
	"myTaskManager/package/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
