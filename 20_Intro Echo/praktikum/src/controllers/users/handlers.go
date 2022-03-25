package users

import (
	"encoding/json"
	"net/http"
	"static-api/src/constants"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"users":   constants.Users,
	})
}

func GetUserController(context echo.Context) error {
	searchUserId, _ := strconv.Atoi(context.Param("id"))
	var userFound constants.User

	for _, user := range constants.Users {
		if user.Id == searchUserId {
			userFound = user
			break
		}
	}

	if userFound == (constants.User{}) {
		return context.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed",
			"result":  "user not found!",
		})
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"uri":     userFound,
	})
}

func CreateUserController(context echo.Context) error {
	jsonBody := constants.User{}
	err := json.NewDecoder(context.Request().Body).Decode(&jsonBody)

	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"result":  "bad request",
		})
	}

	constants.Users = append(constants.Users, jsonBody)

	return context.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"result":  jsonBody,
	})
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

	var userIndex int = -1

	for index, user := range constants.Users {
		if user.Id == userId {
			userIndex = index
			break
		}
	}

	if userIndex == -1 {
		return context.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed",
			"result":  "not found!",
		})
	}

	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"result":  "bad request",
		})
	}

	constants.Users[userIndex] = jsonBody

	return context.JSON(http.StatusNoContent, map[string]interface{}{
		"message": "success",
		"result":  jsonBody,
	})
}

func DeleteUserController(context echo.Context) error {
	var userIndex int = -1
	userId, _ := strconv.Atoi(context.Param("id"))

	for index, user := range constants.Users {
		if user.Id == userId {
			userIndex = index
			break
		}
	}

	if userIndex == -1 {
		return context.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed",
			"result":  "not found!",
		})
	}

	usr := make([]constants.User, 0)
	usr = append(usr, constants.Users[:userIndex]...)
	usr = append(usr, constants.Users[userIndex+1:]...)
	constants.Users = usr

	return context.JSON(http.StatusNoContent, map[string]interface{}{
		"message": "success",
		"result":  "data deleted.",
	})
}
