package userSlug

import (
	"fmt"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create a new user
// @Description Create a new user
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Param client body models.User true "User object"
// @Success 200 {object} map[string]interface{} "User id"
// @Router /user/{user_id} [post]
func (ctrl *Controller) createUser(c *gin.Context) {
	var client models.User

	client.Id, _ = strconv.Atoi(c.Param("user_id"))

	c.BindJSON(&client)

	uuidUser, err := ctrl.service.CreateUser(client)

	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "Failed to create client")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"User id": uuidUser,
	})
}

// @Summary Delete a user
// @Description Delete a user
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Success 200 {object} pkg.StatusResponse "OK"
// @Router /user/{user_id} [delete]
func (ctrl *Controller) deleteUser(c *gin.Context) {
	var client models.User

	var err error
	client.Id, err = strconv.Atoi(c.Param("user_id"))
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "What user_id?")
		return
	}

	if err := ctrl.service.DeleteUser(client.Id); err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "Failed to delete client")
		return
	}

	c.JSON(http.StatusOK, pkg.StatusResponse{
		Status: "OK",
	})
}

// @Summary Get user's slugs
// @Description Get user's slugs
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Success 200 {object} map[string]interface{} "User slugs"
// @Router /user/{user_id} [get]
func (ctrl *Controller) getUser(c *gin.Context) {
	clientId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed read user_id: %s", err.Error()))
		return
	}

	client, err := ctrl.service.GetUser(clientId)
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "Failed find user")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		strconv.Itoa(clientId): client.Slugs,
	})
}

// @Summary Add or delete slugs for a user
// @Description Add or delete slugs for a user
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Param listAddDel body models.AddRemoveUserSlug true "List of slugs to add or delete"
// @Success 200 {object} pkg.StatusResponse "OK"
// @Router /user/{user_id}/slug [post]
func (ctrl *Controller) addDelSlugInUser(c *gin.Context) {
	var listAddDel models.AddRemoveUserSlug
	var client models.User
	var err error = nil

	client.Id, err = strconv.Atoi(c.Param("user_id"))
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Failed read user_id: %s", err.Error()))
		return
	}

	if err = c.BindJSON(&listAddDel); err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = ctrl.service.AddDelSlugToUser(client.Id, listAddDel)
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error with add / del slugs: %s", err.Error()))

	}

	c.JSON(http.StatusOK, pkg.StatusResponse{
		Status: "Ok",
	})
}
