package main

import (
	"Praktikum/configs"
	"Praktikum/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	configs.LoadEnv()
	configs.InitDB()
	routes.InitUserRoutes(e)
	routes.InitBookRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
