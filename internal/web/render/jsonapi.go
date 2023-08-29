package render

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"net/http"
)

const (
	// ContentTypeHeader is a name of header for Content-Type value
	ContentTypeHeader = "Content-Type"
)

// JSONAPIPayload is marshalling function for JSONAPI payload
func JSONAPIPayload(ctx *gin.Context, statusCode int, payload interface{}) {
	ctx.Header(ContentTypeHeader, jsonapi.MediaType)
	ctx.Status(statusCode)

	if err := jsonapi.MarshalPayload(ctx.Writer, payload); err != nil {
		logger.Error("jsonapi.MarshalPayload failed", "error", err)

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
