package handler

import (
	"github.com/AssylzhanZharzhanov/connect/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	return router
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}