package routers

import (
	ctr "rest-api-echo-go/src/controllers"
	"github.com/labstack/echo/v4"
)

func GeneralRouter(gnl *echo.Group) {
	gnl.OPTIONS("/testing-connect-db", ctr.TestConnectDatabase())
	gnl.POST("/testing-connect-db", ctr.TestConnectDatabase())
}