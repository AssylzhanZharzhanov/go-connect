package v1

import (
	"github.com/AssylzhanZharzhanov/connect/internal/event/domain"

	"github.com/gin-gonic/gin"
)

// Handler - represents handler
type Handler struct {
	useCase domain.EventUseCase
}

// NewHandler - creates a new handler with necessary dependencies
func NewHandler(useCase domain.EventUseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}

// Create - Impl.
func (h *Handler) Create(c *gin.Context) {

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
