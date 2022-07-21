package authentications

import (
	"log"
	"middleware-api/src/database"
	"middleware-api/src/database/factories"
	"middleware-api/src/models/auth"
	"middleware-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

var db = database.Connect()
var factory = factories.LoginFactory(db)

func LoginHandler(echoContext echo.Context) error {
	var loginData auth.Login
	echoContext.Bind(&loginData)

	userID, err := factory.AttemptLogin(loginData)
	log.Println(userID, err)
	if err != nil || userID == "" {
		log.Println("Unidentified authentication trials", err)
		return utils.CreateResponse(echoContext, http.StatusUnauthorized, "Can't verify authentication trial.")
	}

	token, _ := utils.GenerateJwt(userID)
	utils.SetAuthCookie(echoContext, token)
	return utils.CreateResponse(echoContext, http.StatusOK, "OK", token)
}
