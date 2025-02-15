package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"todo_app"
	"todo_app/pkg/handler"
	"todo_app/pkg/repository"
	"todo_app/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//	@title			Todo List App
//	@version		1.0
//	@description	REST API Server for TodoList Application

//	@host		localhost:8000
//	@BasePath	/

//	@securitydefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	if err := initConfig(); err != nil{
		logrus.Fatalf("error intitializing config: %v", err)
	}

	if err := godotenv.Load(); err != nil{
		logrus.Fatalf("error loading env variables: %v", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil{
		logrus.Fatalf("error initializing DB: %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func(){
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil{
			logrus.Fatalf("error occured while running http server: %v", err)
		}
	}()

	logrus.Print("TodoApp Started")
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil{
		logrus.Errorf("error occured on server shutting down: %v", err)
	}

	if err := db.Close(); err != nil{
		logrus.Errorf("error occured on db connection close: %v", err)
	}
}

func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}