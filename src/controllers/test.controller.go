package controllers

import (
	"net/http"
	// "time"
	
	util "rest-api-echo-go/src/utils"
	mdl "rest-api-echo-go/src/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	// "github.com/dgrijalva/jwt-go"
	// "github.com/spf13/viper"
)

type ReqParam struct {
	Name string `json:"name" validate:"required"`
	Age string `json:"age" validate:"required"`
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

		//validation request
		if err := ctx.Validate(reqBody); err != nil {
			return ctx.JSON(http.StatusBadRequest, util.CustomResponses{
				"status":  400,
				"message": err.Error(),
			})
		}

		log.Printf("------ data request ---- : %v", reqBody)
		return ctx.JSON(http.StatusOK, reqBody)
	}
}

// type ReqToken struct {
// 	ClientId string `json:"client_id"`
// 	SecretKey string `json:"secret_key"`
// 	GrantType string `json:"grant_type"`
// }

// type jwtCustomClaims struct {
// 	Name  string `json:"name"`
// 	UUID  string `json:"uuid"`
// 	Admin bool   `json:"admin"`
// 	Scopes []string `json:"scopes"`
// 	jwt.StandardClaims
// }

// func GetToken() echo.HandlerFunc {
// 	return func (ctx echo.Context) error {
// 		reqBody := &ReqToken{}
// 		if err := ctx.Bind(&reqBody); err != nil {
// 			return ctx.JSON(http.StatusUnauthorized, util.CustomResponses{
// 				"status": 401,
// 				"message": "Unauthorized. Invalid Parameter.",
// 			})
// 		}

// 		if reqBody.ClientId != "apiv1" || reqBody.SecretKey != "5ae6ea9d886dfb01ca99b8aae3db70d" || reqBody.GrantType != "credentials" {
// 			return ctx.JSON(http.StatusUnauthorized, util.CustomResponses{
// 				"status": 401,
// 				"message": "Unauthorized. Invalid Parameter.",
// 			})
// 		}

// 		exp := time.Now().Add(time.Second * time.Duration(3600)).Unix()

// 		claims := &jwtCustomClaims{
// 			Name:  reqBody.ClientId + reqBody.SecretKey,
// 			UUID:  "123123",
// 			Admin: true,
// 			StandardClaims: jwt.StandardClaims{
// 					ExpiresAt: exp,
// 			},
// 		}

// 		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 		t, err := token.SignedString([]byte(viper.GetString("jwt_secret_key")))
// 		if err != nil {
// 				return err
// 		}

// 		log.Printf("------ data request ---- : %v", reqBody)
// 		return ctx.JSON(http.StatusOK, util.CustomResponses{
// 			"access_token": t,
// 			"expires_in": exp,
// 			"token_type": "Bearer",
// 		})
// 	}
// }


type ReqParamDb struct {
	ClientId string `json:"client_id" validate:"required"`
}

func TestConnectDatabase() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		reqBody := &ReqParamDb{}
		if err := ctx.Bind(&reqBody); err != nil {
			return ctx.JSON(http.StatusUnauthorized, util.CustomResponses{
				"status": 401,
				"message": "Unauthorized. Invalid Parameter.",
			})
		}

		//validation request
		if err := ctx.Validate(reqBody); err != nil {
			return ctx.JSON(http.StatusBadRequest, util.CustomResponses{
				"status":  400,
				"message": err.Error(),
			})
		}

		log.Printf("------ data request ---- : %v", reqBody)

		q := mdl.New(util.DB)

		rows, err := q.GetData(ctx.Request().Context())

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, util.CustomResponses{
				"status":  500,
				"message": err.Error(),
			})
		}

		log.Printf("------ result query ---- : %v", rows)

		return ctx.JSON(http.StatusOK, rows)
	}
}