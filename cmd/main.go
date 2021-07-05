package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/tatatia/list-app-go"
	"github.com/tatatia/list-app-go/pkg/handler"
	"github.com/tatatia/list-app-go/pkg/repository"
	"github.com/tatatia/list-app-go/pkg/service"
	"log"
)

// docker run --name=postgres-db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
//docker start postgres-db//запустити базу після перезавантаження компютера
//docker ps//перевірка таблиць в базі

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while renning http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
