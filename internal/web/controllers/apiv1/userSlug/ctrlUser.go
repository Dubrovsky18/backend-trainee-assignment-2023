package userSlug

import (
	"fmt"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (ctrl *Controller) createUser(c *gin.Context) {
	var client models.User

	if err := c.BindJSON(&client); err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	uuidUser, err := ctrl.service.CreateUser(client)

	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "Failed to create client")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"User id": uuidUser,
	})
}

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

func (ctrl *Controller) getUser(c *gin.Context) {
	var client models.User

	var err error
	client.Id, err = strconv.Atoi(c.Param("user_id"))
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "Failed read user_id")
		return
	}

	client, err = ctrl.service.GetUser(client.Id)
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "Failed find user")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		fmt.Sprintf("User: %s", client.Id): fmt.Sprintf("Slugs: %s", client.Slugs),
	})
}

func (ctrl *Controller) AddDelSlugInUser(c *gin.Context) {
	var listAddDel models.AddRemoveUserSlug
	var client models.User
	var err error

	client.Id, err = strconv.Atoi(c.Param("user_id"))
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, "Failed read user_id")
		return
	}

	if err = c.BindJSON(&listAddDel); err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = ctrl.service.AddDelSlugToUser(client.Id, listAddDel)
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to add or del for user: %s"))
		return
	}

	c.JSON(http.StatusOK, pkg.StatusResponse{
		Status: "Ok",
	})

}
