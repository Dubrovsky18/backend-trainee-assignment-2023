package userSlug

import (
	services "github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/services"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/web/controllers/apiv1"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	apiv1.BaseController
	service services.TemplateServiceUserSlug
}

func NewController(service services.TemplateServiceUserSlug) *Controller {
	return &Controller{
		service: service,
	}
}

// DefineRoutes adds controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {

	apiVer := r.Group("/api/v1")
	users := apiVer.Group("/users")

	{
		users.POST("/create", ctrl.createUser)
		users.POST("/add_del_slug/:user_id", ctrl.AddDelSlugInUser)
		users.GET("/get_slugs/:user_id", ctrl.getUser)
		users.DELETE("/delete/:user_id", ctrl.deleteUser)

	}

	slug := apiVer.Group("/slug")

	{
		slug.POST("/create", ctrl.createSlug)
		slug.DELETE("/delete/:name_slug", ctrl.deleteSlug)

	}
}
