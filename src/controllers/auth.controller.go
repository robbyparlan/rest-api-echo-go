package controllers

import (
	"net/http"
	"time"
	
	util "rest-api-echo-go/src/utils"
	mdl "rest-api-echo-go/src/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type ParamsApiClient struct {
	ClientId string `json:"client_id" validate:"required"`
	SecretKey string `json:"secret_key" validate:"required"`
	GrantType string `json:"grant_type" validate:"required"`
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
	Admin bool   `json:"admin"`
	Scopes []string `json:"scopes"`
	jwt.StandardClaims
}

func GetTokenV1() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		reqBody := &ParamsApiClient{}
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

		q := mdl.New(util.DB)

		rows, err := q.GetDataApiClient(ctx.Request().Context(),
			mdl.ParamApiClient{
				ClientId: reqBody.ClientId,
				SecretKey: reqBody.SecretKey,
			},
		)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, util.CustomResponses{
				"status":  500,
				"message": err.Error(),
			})
		}

		log.Printf("------ result query ---- : %v", rows)

		exp := time.Now().Add(time.Second * time.Duration(3600)).Unix()

		claims := &jwtCustomClaims{
			Name:  rows.ClientId + rows.SecretKey,
			UUID:  rows.Uuid,
			Admin: true,
			Scopes: rows.Scopes,
			StandardClaims: jwt.StandardClaims{
					ExpiresAt: exp,
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte(viper.GetString("jwt_secret_key")))
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
