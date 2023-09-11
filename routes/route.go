package routes

import (
	"goLANG/handlers"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	authGroup := e.Group("")
	authGroup.Use(handlers.TokenMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/register", handlers.CreateUser)
	e.GET("/register", handlers.GetUser)
	authGroup.PUT("/register/:id", handlers.UpdateUser)
	authGroup.DELETE("/register/:id", handlers.DeleteUser)

	e.POST("/login", handlers.LoginAccount)

	return e
}