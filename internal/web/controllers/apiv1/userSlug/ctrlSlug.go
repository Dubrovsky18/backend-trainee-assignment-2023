package userSlug

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl *Controller) createSlug(c *gin.Context) {
	var slug models.Slug

	if err := c.BindJSON(&slug); err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.service.CreateSlug(slug); err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "Failed to create slug")
		return
	}

	c.JSON(http.StatusOK, pkg.StatusResponse{
		Status: "OK",
	})

}

func (ctrl *Controller) deleteSlug(c *gin.Context) {
	var slug models.Slug

	slug.NameSlug = c.Param("name_slug")

	if err := ctrl.service.DeleteSlug(slug); err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "Failed to delete slug")
		return
	}

	c.JSON(http.StatusOK, pkg.StatusResponse{
		Status: "OK",
	})

}

func (ctrl *Controller) getSlugs(c *gin.Context) {
	result, err := ctrl.service.GetSlugs()
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
