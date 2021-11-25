package routers

import (
	ctr "rest-api-echo-go/src/controllers"
	"github.com/labstack/echo/v4"
)

func AuthRouter(auth *echo.Group) {
	auth.OPTIONS("/access_token", ctr.GetToken())
	auth.POST("/access_token", ctr.GetToken())
}