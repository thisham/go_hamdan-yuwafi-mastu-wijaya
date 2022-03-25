package users

import (
	"encoding/json"
	"net/http"
	"static-api/src/constants"
	"static-api/src/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(context echo.Context) error {
	return utils.CreateResponse(context, http.StatusOK, "got users list", constants.Users)
}

func GetUserController(context echo.Context) error {
	userId, _ := strconv.Atoi(context.Param("id"))
	var foundUserIndex = findUserIndex(userId)

	if foundUserIndex == -1 {
		return utils.CreateResponse(context, http.StatusNotFound, "user not found!")
	}

	return utils.CreateResponse(context, http.StatusOK, "user found!", constants.Users[foundUserIndex])
}

func CreateUserController(context echo.Context) error {
	jsonBody := constants.User{}
	err := json.NewDecoder(context.Request().Body).Decode(&jsonBody)

	if err != nil {
		return utils.CreateResponse(context, http.StatusBadRequest, "bad request recived.")
	}

	constants.Users = append(constants.Users, jsonBody)
	return utils.CreateResponse(context, http.StatusCreated, "user created.")
}

func UpdateUserController(context echo.Context) error {
	userId, _ := strconv.Atoi(context.Param("id"))
	jsonBody := constants.User{
		Id:       userId,
		Name:     "",
		Email:    "",
		Password: "",
	}
	err := json.NewDecoder(context.Request().Body).Decode(&jsonBody)

	if err != nil {
		return utils.CreateResponse(context, http.StatusBadRequest, "bad request received.")
	}

	var foundUserIndex = findUserIndex(userId)

	if foundUserIndex == -1 {
		return utils.CreateResponse(context, http.StatusNotFound, "user not found.")
	}

	constants.Users[foundUserIndex] = jsonBody
	return utils.CreateResponse(context, http.StatusNoContent, "user data updated.")
}

func DeleteUserController(context echo.Context) error {
	userId, _ := strconv.Atoi(context.Param("id"))
	var foundUserIndex = findUserIndex(userId)

	if foundUserIndex == -1 {
		return utils.CreateResponse(context, http.StatusNotFound, "user not found!")
	}

	usr := make([]constants.User, 0)
	usr = append(usr, constants.Users[:foundUserIndex]...)
	usr = append(usr, constants.Users[foundUserIndex+1:]...)
	constants.Users = usr

	return utils.CreateResponse(context, http.StatusNoContent, "data updated.")
}
