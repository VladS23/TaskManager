package main

import (
	"github.com/spf13/viper"
	"log"
	todo "myTaskManager"
	"myTaskManager/package/handler"
	"myTaskManager/package/repository"
	"myTaskManager/package/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config %s", err.Error())
	}
	repos := repository.NewRepositoty()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
