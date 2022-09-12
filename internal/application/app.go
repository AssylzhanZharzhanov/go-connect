package application

import (
	"context"
	"github.com/AssylzhanZharzhanov/connect/internal/config"
	"net/http"
	"time"
)

type App struct {
	cfg        config.Config
	httpServer *http.Server
}

func NewApp(cfg config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (s *App) Run() error {

	// initialize tracer
	// initialize logger
	// initialize kafka connection
	// initialize db connection
	// initialize repository
	// initialize usecase
	// initialize controllers or handlers

	s.httpServer = &http.Server{
		Addr: ":" + s.cfg.Port,
		//Handler:        s.Handlers,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()

	// initialize kafka consumer
}

func (s *App) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
