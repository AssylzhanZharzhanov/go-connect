package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AssylzhanZharzhanov/connect/internal/application"
	"github.com/AssylzhanZharzhanov/connect/pkg/utils"
	"github.com/sirupsen/logrus"
)

func main() {

	var (
		ctx = context.Background()
	)

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

	if err := app.Shutdown(ctx); err != nil {
		logrus.Errorf("error occured on application shutting down: %s", err.Error())
	}
}
