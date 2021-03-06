package routers

import (
	ctr "rest-api-echo-go/src/controllers"
	"github.com/labstack/echo/v4"
)

func TestRouter(api *echo.Group) {
	api.OPTIONS("/test", ctr.Test())
	api.POST("/test", ctr.Test())

	api.OPTIONS("/test-db", ctr.TestConnectDatabase())
	api.POST("/test-db", ctr.TestConnectDatabase())
}