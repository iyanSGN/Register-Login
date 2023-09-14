package main

import (
	"goLANG/pkg/database"
	"goLANG/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	database.InitDB()
	database.Migrate()

	routes.Init(e)

    e.Logger.Fatal(e.Start(":8090"))
}