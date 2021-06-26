package main

import (
	"github.com/tatatia/list-app-go"
	"github.com/tatatia/list-app-go/pkg/handler"
	"github.com/tatatia/list-app-go/pkg/repository"
	"github.com/tatatia/list-app-go/pkg/service"
	//"github.com/zhashkevych/todo-app"
	//"github.com/zhashkevych/todo-app/pkg/handler"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("3306", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while renning http server: %s", err.Error())
	}
}
