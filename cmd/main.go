package main

import (
	"log"
	todo "myTaskManager"
	"myTaskManager/package/handler"
	"myTaskManager/package/repository"
	"myTaskManager/package/service"
)

func main() {
	repos := repository.NewRepositoty()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
