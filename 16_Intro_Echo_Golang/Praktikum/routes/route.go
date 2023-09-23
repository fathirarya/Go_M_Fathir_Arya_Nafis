package routes

import (
	"Praktikum/controller"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	e.GET("/users", controller.GetUsersControllers)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
}
