package application

import (
	"context"
	"github.com/AssylzhanZharzhanov/connect/pkg/jaeger"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"time"

	"github.com/AssylzhanZharzhanov/connect/internal/config"
	eventHandlerV1 "github.com/AssylzhanZharzhanov/connect/internal/event/delivery/http/v1"
	eventRepository "github.com/AssylzhanZharzhanov/connect/internal/event/repository"
	eventUseCase "github.com/AssylzhanZharzhanov/connect/internal/event/usecase"
	"github.com/AssylzhanZharzhanov/connect/pkg/db/postgres"

	"github.com/gin-gonic/gin"
)

// App - stores all configuration of the application.
type App struct {
	cfg        config.AppConfig
	httpServer *http.Server
}

// NewApp - creates a new application
func NewApp(cfg config.AppConfig) *App {
	return &App{
		cfg: cfg,
	}
}

// Run - starts application
func (s *App) Run() error {

	router := gin.Default()
	//router.Use(
	//	gin.Recovery(),
	//	gin.Logger(),
	//)

	api := router.Group("/api")

	// initialize logger
	// initialize kafka connection

	// enable tracing
	if s.cfg.Enable {
		tracer, closer, err := jaeger.NewJaegerTracer(s.cfg)
		if err != nil {
			return err
		}
		defer closer.Close()
		opentracing.SetGlobalTracer(tracer)
	}

	//db connection
	db, err := postgres.NewPostgresDB(s.cfg.DSN)
	if err != nil {
		return err
	}

	//Repository layer
	eventsRepository := eventRepository.NewRepository(db)

	//Service layer
	eventsUseCase := eventUseCase.NewUseCase(eventsRepository)

	//Handler layer
	eventHandlerV1.RegisterEndpoints(api, eventsUseCase)

	s.httpServer = &http.Server{
		Addr:           ":" + s.cfg.Port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()

	// initialize kafka consumer
}

// Shutdown - stops application
func (s *App) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
