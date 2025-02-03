package main

import (
	"github.com/beebablaster/todo-app"
	"github.com/beebablaster/todo-app/pkg/handler"
	"github.com/beebablaster/todo-app/pkg/repository"
	"github.com/beebablaster/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
