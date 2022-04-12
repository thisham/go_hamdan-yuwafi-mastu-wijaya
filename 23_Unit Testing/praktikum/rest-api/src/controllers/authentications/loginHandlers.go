package authentications

import (
	"log"
	"middleware-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginHandler(echoContext echo.Context) error {
	token, err := utils.GenerateJwt()

	if err != nil {
		log.Println("Error when creating JWT token", err)
		return utils.CreateResponse(echoContext, http.StatusInternalServerError, "ERROR: something went wrong while creating JWT token!")
	}

	utils.SetAuthCookie(echoContext, token)
	return utils.CreateResponse(echoContext, http.StatusOK, "OK", token)
}
