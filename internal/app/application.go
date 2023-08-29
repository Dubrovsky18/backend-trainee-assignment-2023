package app

import (
	"context"
	"fmt"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/dependencies"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/initializers"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/repository"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/services"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/pkg/logger"
	"net/http"
)

// Application is a main struct for the application that contains general information
type Application struct {
	httpServer *http.Server
	Container  *dependencies.Container
}

// InitializeApplication initializes new application
func InitializeApplication() (*Application, error) {
	initializers.InitializeEnv()
	appConfig := initializers.InitializeAppConfig()
	info := initializers.InitializeBuildInfo()

	templateUserSlugRepo := repository.NewTemplateRepository()
	serviceUserSlug := services.NewTemplateRepository(templateUserSlugRepo)

	container := &dependencies.Container{
		BuildInfo: info,
		Service:   serviceUserSlug,
	}

	router := initializers.InitializeRouter(container)
	server := initializers.InitializeHTTPServer(router, appConfig.HTTP)

	return &Application{
		httpServer: server,
		Container:  container,
	}, nil
}

// Start starts application services
func (a *Application) Start(ctx context.Context, cli bool) {
	if cli {
		return
	}

	a.startHTTPServer()
}

// Stop stops application services
func (a *Application) Stop() (err error) {
	return a.httpServer.Shutdown(context.TODO())
}

func (a *Application) startHTTPServer() {
	go func() {
		logger.Info(fmt.Sprintf("started http server on address: %s", a.httpServer.Addr))

		// service connections
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("HTTP Server stopped", "error", err)
		}
	}()
}
