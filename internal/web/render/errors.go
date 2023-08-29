package render

import (
	"encoding/json"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"github.com/google/jsonapi"
	"net/http"
)

type ResponseError struct {
	Error       string `json:"error"`
	Message     string `json:"message"`
	ServiceName string `json:"serviceName"`
}

var ServiceName string = envy.Get("SERVICE_NAME", "Gobase")

func SendError(ctx *gin.Context, statusCode int, err error, message string) {

	ctx.Header(ContentTypeHeader, jsonapi.MediaType)
	ctx.Status(statusCode)
	var response = ResponseError{
		Error:       err.Error(),
		Message:     message,
		ServiceName: ServiceName,
	}

	if err := json.NewEncoder(ctx.Writer).Encode(&response); err != nil {
		logger.Error("jsonapi.MarshalErrors failed", "error", err)

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
