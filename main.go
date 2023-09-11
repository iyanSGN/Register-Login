package main

import (
	"goLANG/pkg/database"
	"goLANG/routes"
)

func main() {

	database.InitDB()
	database.Migrate()

	e := routes.Init()

    e.Logger.Fatal(e.Start(":8090"))
}