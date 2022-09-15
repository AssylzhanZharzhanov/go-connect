package v1

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/AssylzhanZharzhanov/connect/internal/event/domain"
	"github.com/AssylzhanZharzhanov/connect/pkg/errors"
	"github.com/AssylzhanZharzhanov/connect/pkg/jaeger"

	"github.com/gin-gonic/gin"
)

// Handler - represents handler
type Handler struct {
	useCase domain.EventUseCase
	logger  *zap.SugaredLogger
}

// NewHandler - creates a new handler with necessary dependencies
func NewHandler(useCase domain.EventUseCase, logger *zap.SugaredLogger) Handler {
	return Handler{
		useCase: useCase,
		logger:  logger,
	}
}

// Create - Impl.
func (h *Handler) Create(c *gin.Context) {
	ctx, span := jaeger.StartHTTPServerTracerSpan(c, "event.Create")
	defer span.Finish()

	var input domain.CreateEventRequestDTO

	if err := c.BindJSON(&input); err != nil {
		h.logger.Infof("events.Create err: %v", jaeger.TraceWithErr(span, err))
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.useCase.CreateEvent(ctx, input)
	if err != nil {
		h.logger.Infof("events.Create err: %v", jaeger.TraceWithErr(span, err))
		errors.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "")
}

// List - Impl.
func (h *Handler) List(c *gin.Context) {

}

// Get - Impl.
func (h *Handler) Get(c *gin.Context) {

}

// Update - Impl.
func (h *Handler) Update(c *gin.Context) {

}

// Delete - Impl.
func (h *Handler) Delete(c *gin.Context) {

}
