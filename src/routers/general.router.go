package routers

import (
	ctr "rest-api-echo-go/src/controllers"
	"github.com/labstack/echo/v4"
)

func GeneralRouter(gnl *echo.Group) {
	gnl.OPTIONS("/test", ctr.TestConnectDatabase())
	gnl.POST("/test", ctr.TestConnectDatabase())
}