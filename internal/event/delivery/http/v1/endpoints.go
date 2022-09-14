package v1

import (
	"github.com/AssylzhanZharzhanov/connect/internal/event/domain"

	"github.com/gin-gonic/gin"
)

// RegisterEndpoints - registers new endpoints
func RegisterEndpoints(router *gin.RouterGroup, useCase domain.EventUseCase) {
	h := NewHandler(useCase)

	users := router.Group("/events")
	{
		users.POST("/", h.Create)
		users.GET("/", h.List)
		users.GET("/:id", h.Get)
		users.PUT("/:id", h.Update)
		users.DELETE("/:id", h.Delete)
	}
}
