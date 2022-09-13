package main

import (
	"context"
	"github.com/AssylzhanZharzhanov/connect/pkg/utils"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AssylzhanZharzhanov/connect/internal/application"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg, err := utils.LoadConfig()
	if err != nil {
		return
	}

	app := application.NewApp(cfg)
	go func() {
		if err := app.Run(); err != nil {
			log.Fatalf("Error in starting application: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := app.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on application shutting down: %s", err.Error())
	}
}
