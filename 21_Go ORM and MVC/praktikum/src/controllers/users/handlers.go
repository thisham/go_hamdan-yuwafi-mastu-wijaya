package users

import (
	"gorm-mvc-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(ctx echo.Context) error {
	return utils.CreateResponse(ctx, http.StatusOK, "delivered")
}
