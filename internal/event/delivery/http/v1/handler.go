package v1

import (
	"github.com/AssylzhanZharzhanov/connect/internal/event/domain"
	"github.com/gin-gonic/gin"
)

type handler struct {
	useCase domain.EventUseCase
}

func NewHandler(useCase domain.EventUseCase) *handler {
	return &handler{}
}

func (h *handler) Create(c *gin.Context) {

}

func (h *handler) List(c *gin.Context) {

}

func (h *handler) Get(c *gin.Context) {

}

func (h *handler) Update(c *gin.Context) {

}

func (h *handler) Delete(c *gin.Context) {

}
