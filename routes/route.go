package routes

import (
	"goLANG/handlers"
	"net/http"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	authGroup := e.Group("")
	authGroup.Use(handlers.TokenMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/register", handlers.CreateUser)
	e.GET("/register", handlers.GetUser)
	e.GET("/register/:id", handlers.GetId)
	authGroup.PUT("/register/:id", handlers.UpdateUser)
	authGroup.DELETE("/register/:id", handlers.DeleteUser)


	e.GET("/department_id/:department_id", handlers.GetUserByDepId)
	e.POST("/login", handlers.LoginAccount)
}