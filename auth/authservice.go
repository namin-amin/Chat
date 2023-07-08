package auth

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTAuth
//
// authenticate the received jwt token
func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sentToken := c.Request().Header.Values("Authorization")

		if len(sentToken) == 0 {
			fmt.Println("no token")
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "no auth token"})
		}
		token, err := jwt.ParseWithClaims(sentToken[0], &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			fmt.Println("jwt issue")
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "malformed jwt"})
		}

		claims := token.Claims.(*JwtCustomClaims)

		// c.Set("name", claims.Name)
		// c.Set("admin", claims.Admin)
		fmt.Println(claims.ID)
		c.Request().Header.Set("userId", claims.ID)
		return next(c)
	}
}
