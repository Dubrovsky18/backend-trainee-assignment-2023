package pkg

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ErrorResponse struct {
	Message string `json:*message*`
}

type StatusResponse struct {
	Status string `json:*Status`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Fatalf("Error: %s", message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})

}
