package main

import (
	"context"
	"github.com/AssylzhanZharzhanov/connect/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AssylzhanZharzhanov/connect/internal/application"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error in reading env file: %s", err.Error())
	}
}

func main() {

	//load config

	restSrv := application.NewApp(config.Config{})
	go func() {
		if err := restSrv.Run(); err != nil {
			log.Fatalf("Error in starting application: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := restSrv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on application shutting down: %s", err.Error())
	}
}
