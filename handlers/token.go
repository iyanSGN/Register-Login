package handlers

import (
	"goLANG/token"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return c.JSON(http.StatusUnauthorized, map[string]interface{}{
                "error": "Token is missing",
				"status_code": http.StatusUnauthorized,
            })
        }

        tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

        newToken, err := token.VerifyToken(tokenString)
        if err != nil {
            return c.JSON(http.StatusUnauthorized, map[string]interface{}{
                "error": "Invalid token",
				"status_code": http.StatusUnauthorized,
            })
        }

        if newToken != tokenString {
            c.Set("newToken", newToken)
        }

        return next(c)
    }
}
