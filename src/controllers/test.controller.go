package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	util "rest-api-echo-go/src/utils"
	"github.com/labstack/gommon/log"
)

type ReqParam struct {
	Name string `json:"name"`
	Age string `json:"age"`
}

func Test() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		reqBody := &ReqParam{}
		if err := ctx.Bind(&reqBody); err != nil {
			return ctx.JSON(http.StatusUnauthorized, util.CustomResponses{
				"status": 401,
				"message": "Unauthorized. Invalid Parameter.",
			})
		}
		log.Printf("------ data request ---- : %v", reqBody)
		return ctx.JSON(http.StatusOK, reqBody)
	}
}