package users

import (
	"gorm-mvc-api/src/database"
	"gorm-mvc-api/src/database/factories"
	"gorm-mvc-api/src/models/user"
	"gorm-mvc-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

var db = database.Connect()
var factory = factories.UserFactory(db)

func GetAllUsers(echoContext echo.Context) error {
	users, err := factory.GetAllUsers()

	if err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed", err)
	}

	return utils.CreateResponse(echoContext, http.StatusOK, "OK", users)
}

func GetUser(echoContext echo.Context) error {
	var userID string = echoContext.Param("id")
	userData, err := factory.GetUser(userID)

	if err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed", err)
	}

	return utils.CreateResponse(echoContext, http.StatusOK, "OK", userData)
}

func CreateUser(echoContext echo.Context) error {
	var request user.User
	echoContext.Bind(&request)

	if err := factory.CreateUser(request); err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed")
	}

	return utils.CreateResponse(echoContext, http.StatusCreated, "user created")
}

func UpdateUser(echoContext echo.Context) error {
	var request user.User
	echoContext.Bind(&request)

	if err := factory.UpdateUser(request, echoContext.Param("id")); err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed", err)
	}

	return utils.CreateResponse(echoContext, http.StatusNoContent, "user updated")
}
