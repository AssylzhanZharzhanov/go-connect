package application

import (
	"context"
	"net/http"
	"time"

	"github.com/AssylzhanZharzhanov/connect/internal/config"
	eventRepository "github.com/AssylzhanZharzhanov/connect/internal/event/repository"
	eventUseCase "github.com/AssylzhanZharzhanov/connect/internal/event/usecase"
	"github.com/AssylzhanZharzhanov/connect/pkg/db/postgres"
)

type App struct {
	cfg        config.AppConfig
	httpServer *http.Server
}

func NewApp(cfg config.AppConfig) *App {
	return &App{
		cfg: cfg,
	}
}

func (s *App) Run() error {

	// initialize tracer
	// initialize logger
	// initialize kafka connection

	//db connection
	db, err := postgres.NewPostgresDB(s.cfg.DSN)
	if err != nil {
		return err
	}

	//Repository layer
	eventsRepository := eventRepository.NewRepository(db)

	//Service layer
	_ = eventUseCase.NewUseCase(eventsRepository)
	//Handler layer

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
