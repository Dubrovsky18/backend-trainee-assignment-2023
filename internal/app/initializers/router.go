package initializers

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/dependencies"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/web/controllers/apiv1"
	apiv1Status "github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/web/controllers/apiv1/status"
	apiv1SericeUserSlug "github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/web/controllers/apiv1/userSlug"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/web/router"
	"github.com/gin-gonic/gin"
)

// InitializeRouter initializes new gin router
func InitializeRouter(container *dependencies.Container) *gin.Engine {
	r := router.NewRouter()

	ctrls := buildControllers(container)

	for i := range ctrls {
		ctrls[i].DefineRoutes(r)
	}

	return r
}

func buildControllers(container *dependencies.Container) []apiv1.Controller {
	return []apiv1.Controller{
		apiv1Status.NewController(container.BuildInfo),
		apiv1SericeUserSlug.NewController(container.Service),
	}
}
