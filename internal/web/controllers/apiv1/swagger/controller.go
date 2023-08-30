package swagger

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/web/controllers/apiv1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	//nolint: golint //reason: blank import because of swagger docs init
	_ "github.com/Dubrovsky18/backend-trainee-assignment-2023/api"
)

var (
	_ apiv1.Controller = (*Controller)(nil)
)

// Controller implements controller for swagger
type Controller struct {
	apiv1.BaseController
}

// NewController create new instance for swagger controller
func NewController() *Controller {
	return &Controller{}
}

// DefineRoutes adds swagger controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
