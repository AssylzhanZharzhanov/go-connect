package main

import (
	"context"
	"github.com/AssylzhanZharzhanov/connect/pkg/handler"
	"github.com/AssylzhanZharzhanov/connect/pkg/repository"
	"github.com/AssylzhanZharzhanov/connect/pkg/service"
	"github.com/AssylzhanZharzhanov/connect/server"
	"github.com/AssylzhanZharzhanov/connect/utils/db/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error in reading env file: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DBName: viper.GetString("database.name"),
		SSLMode: viper.GetString("database.ssl_mode"),
	})
	if err != nil {
		log.Fatalf("Error in starting database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	port := viper.GetString("server.port")
	restSrv := server.NewServer(handlers.InitRoutes(), port)

	go func() {
		if err := restSrv.Run(); err != nil {
			log.Fatalf("Error in starting server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := restSrv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
