package users

import (
	"gorm-mvc-api/src/database"
	"gorm-mvc-api/src/database/factories"
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
