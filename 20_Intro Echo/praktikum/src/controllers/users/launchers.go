package users

import (
	"net/http"
	"static-api/src/constants"
	"static-api/src/utils"

	"github.com/labstack/echo/v4"
)

func createUser(context echo.Context, userData constants.User) error {
	constants.Users = append(constants.Users, userData)
	return utils.CreateResponse(context, http.StatusCreated, "user created.")
}

func updateUser(context echo.Context, userIndex int, userId int, requestBody constants.User) error {
	constants.Users[userIndex] = requestBody
	return utils.CreateResponse(context, http.StatusNoContent, "user data updated.")
}

func deleteUser(context echo.Context, userIndex int) error {
	partialUsers := make([]constants.User, 0)
	partialUsers = append(partialUsers, constants.Users[:userIndex]...)
	partialUsers = append(partialUsers, constants.Users[userIndex+1:]...)
	constants.Users = partialUsers

	return utils.CreateResponse(context, http.StatusNoContent, "data updated.")
}
