package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	cfg "echo-go/src/config"
	"github.com/labstack/gommon/log"
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

	e.Use(middleware.CORSWithConfig(CORSConfig))


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	type ReqParam struct {
		Name string `json:"name"`
		Age string `json:"age"`
	}

	type H map[string]interface{}

	e.POST("/test", func(ctx echo.Context) error {
		reqBody := &ReqParam{}
		if err := ctx.Bind(&reqBody); err != nil {
			return ctx.JSON(http.StatusUnauthorized, H{
				"status": 401,
				"message": "Unauthorized. Invalid Parameter.",
			})
		}
		log.Printf("------ data request ---- : %v", reqBody)
		return ctx.JSON(http.StatusOK, reqBody)
	})
	e.Logger.Fatal(e.Start(":1323"))
}