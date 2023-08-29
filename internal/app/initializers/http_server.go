package initializers

import (
	"fmt"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitializeHTTPServer create new http.Server instance
func InitializeHTTPServer(router *gin.Engine, cfg *config.HTTPConfig) *http.Server {
	// create http server
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler: router,
	}
}
