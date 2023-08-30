package userSlug

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Create a new slug
// @Description Create a new slug
// @Accept  json
// @Produce  json
// @Param slug body models.Slug true "Slug object"
// @Success 200 {object} pkg.StatusResponse "OK"
// @Router /slug [post]
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

// @Summary Delete a slug
// @Description Delete a slug
// @Accept  json
// @Produce  json
// @Param name_slug path string true "Name of the slug"
// @Success 200 {object} pkg.StatusResponse "OK"
// @Router /slug/{name_slug} [delete]
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

// @Summary Get all slugs
// @Description Get all slugs
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Slug "List of slugs"
// @Router /slug [get]
func (ctrl *Controller) getSlugs(c *gin.Context) {
	result, err := ctrl.service.GetSlugs()
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
