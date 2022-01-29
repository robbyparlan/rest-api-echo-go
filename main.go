package main

import (
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	cfg "rest-api-echo-go/src/config"
	mddl "rest-api-echo-go/src/middlewares"

	router "rest-api-echo-go/src/routers"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/locales/id"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	translator "github.com/go-playground/universal-translator"
	"errors"
	"fmt"
)

type CustomValidator struct {
	validator *validator.Validate
}

/*
	@author Roby Parlan
	Function custom validation request body in bahasa
*/
func (cv *CustomValidator) Validate(i interface{}) error {
	id := id.New()
	uni := translator.New(id, id)

	// translate into bahasa
	trans, _ := uni.GetTranslator("id")
	id_translations.RegisterDefaultTranslations(cv.validator, trans)
	
	err := cv.validator.Struct(i)

	if err != nil {
		object, _ := err.(validator.ValidationErrors)

		for _, key := range object {
			return errors.New(key.Translate(trans))
		}
	}

	return nil
}

func main() {
	cfg.InitConfig()

	e := echo.New()
	e.HideBanner = true
	e.Debug = viper.GetBool("debug")

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] : ${status} " + `{ method: "${method}", uri: "${uri}", request_id: "${id}", remote_ip: "${remote_ip}" }` +
		` { host: "${host}", x-real-ip: "${header:X-Real-IP}", latency: "${latency}", user_agent: "${user_agent}", error: "${error}" }` + "\n",
		Output: cfg.LogWriteFile,
	}))	

	e.Validator = &CustomValidator{validator: validator.New()}

	e.HTTPErrorHandler = func(err error, c echo.Context) {
    report, ok := err.(*echo.HTTPError)
    if !ok {
        report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    if castedObject, ok := err.(validator.ValidationErrors); ok {
        for _, err := range castedObject {
            switch err.Tag() {
            case "required":
                report.Message = fmt.Sprintf("%s tidak boleh kosong", 
                    err.Field())
            }

            break
        }
    }

    c.Logger().Error(report)
    c.JSON(report.Code, report)
	}

	e.Use(middleware.Recover())

	//CORS Config
	CORSConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}
	
	auth := e.Group("/auth")
	api := e.Group("/api")
	gnl := e.Group("/apiv1")

	auth.Use(middleware.CORSWithConfig(CORSConfig))
	api.Use(middleware.CORSWithConfig(CORSConfig))
	gnl.Use(middleware.CORSWithConfig(CORSConfig))

	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(viper.GetString("jwt_secret_key")),
		SigningMethod: "HS256",
	}))

	gnl.Use(mddl.HandleApiAdmin)

	router.AuthRouter(auth)
	router.TestRouter(api)

	router.GeneralRouter(gnl)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Bismillah, Restful API Golang")
	})

	server := &http.Server{
		Addr:         ":" + viper.GetString("app_port"),
		ReadTimeout:  time.Duration(viper.GetInt("app_rto")) * time.Second,
		WriteTimeout: time.Duration(viper.GetInt("app_wto")) * time.Second,
	}
	e.Logger.Fatal(e.StartServer(server))
}