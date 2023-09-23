package main

import (
	"Praktikum/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8000"))
}
