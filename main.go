package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	cfg "rest-api-echo-go/src/config"

	router "rest-api-echo-go/src/routers"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] : ${status} " + `{ method: "${method}", uri: "${uri}", request_id: "${id}", remote_ip: "${remote_ip}" }` +
		` { host: "${host}", x-real-ip: "${header:X-Real-IP}", latency: "${latency}", user_agent: "${user_agent}", error: "${error}" }` + "\n",
		Output: cfg.LogWriteFile,
	}))	

	e.Use(middleware.Recover())

	//CORS Config
	CORSConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}
	
	api := e.Group("/api")

	api.Use(middleware.CORSWithConfig(CORSConfig))

	router.TestRouter(api)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Bismillah, Restful API Golang")
	})

	e.Logger.Fatal(e.Start(":1323"))
}