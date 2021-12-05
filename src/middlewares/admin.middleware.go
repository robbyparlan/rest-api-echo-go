package middlewares

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"

	util "rest-api-echo-go/src/utils"
)

func HandleApiAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func (ctx echo.Context) error {
		_, authHead := ctx.Request().Header["Authorization"]

		if !authHead {
			return ctx.JSON(http.StatusUnauthorized, util.CustomResponses{
				"status": 401,
				"message": "Unauthorized. Invalid Parameter.",
			})
		}

		data := ctx.Request().Header.Get("Authorization")

		key := data[0:6]

		if key != "Bearer" {
			return ctx.JSON(http.StatusUnauthorized, util.CustomResponses{
				"status": 401,
				"message": "Unauthorized. Parameter Must Be [Bearer]",
			})
		}

		tokenString := data[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt_secret_key")), nil
		})

		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, util.CustomResponses{
				"status": 401,
				"message": err.Error(),
			})
		}
		
		claims := token.Claims.(jwt.MapClaims)

		scopes := claims["scopes"].([]interface{})

		for _, v := range scopes {
			if v == "ADMIN" {
					break;
				} else {
					return ctx.JSON(http.StatusUnauthorized, util.CustomResponses{
						"status": 401,
						"message": "Unauthorized. Access Denied!",
					})	
			}
		}

		return next(ctx)

	}
}