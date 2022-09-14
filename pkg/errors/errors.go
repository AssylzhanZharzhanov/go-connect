package errors

import "github.com/gin-gonic/gin"

type errors struct {
	Message string
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errors{Message: message})
}
