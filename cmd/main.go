package main

import (
	_ "github.com/jackc/pgx"
	"github.com/spf13/viper"
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
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
