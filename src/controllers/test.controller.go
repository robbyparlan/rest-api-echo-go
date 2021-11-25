package controllers

import (
	"net/http"
	"time"
	
	util "rest-api-echo-go/src/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/dgrijalva/jwt-go"
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

type ReqToken struct {
	ClientId string `json:"client_id"`
	SecretKey string `json:"secret_key"`
	GrantType string `json:"grant_type"`
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func GetToken() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		reqBody := &ReqToken{}
		if err := ctx.Bind(&reqBody); err != nil {
			return ctx.JSON(http.StatusUnauthorized, util.CustomResponses{
				"status": 401,
				"message": "Unauthorized. Invalid Parameter.",
			})
		}

		if reqBody.ClientId != "apiv1" || reqBody.SecretKey != "5ae6ea9d886dfb01ca99b8aae3db70d" || reqBody.GrantType != "credentials" {
			return ctx.JSON(http.StatusUnauthorized, util.CustomResponses{
				"status": 401,
				"message": "Unauthorized. Invalid Parameter.",
			})
		}

		exp := time.Now().Add(time.Second * time.Duration(3600)).Unix()

		claims := &jwtCustomClaims{
			Name:  reqBody.ClientId + reqBody.SecretKey,
			UUID:  "123123",
			Admin: true,
			StandardClaims: jwt.StandardClaims{
					ExpiresAt: exp,
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte(reqBody.SecretKey))
		if err != nil {
				return err
		}

		log.Printf("------ data request ---- : %v", reqBody)
		return ctx.JSON(http.StatusOK, util.CustomResponses{
			"access_token": t,
			"expires_in": exp,
			"token_type": "Bearer",
		})
	}
}