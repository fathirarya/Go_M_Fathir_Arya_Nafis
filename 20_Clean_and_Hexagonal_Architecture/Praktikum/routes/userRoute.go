package routes

import (
	"belajar-go-echo/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RoutesWithoutController(e *echo.Echo, userController controller.UserController) {
	usersGroup := e.Group("users")

	usersGroup.POST("", userController.CreateUser())

}

func RoutesWithMiddleware(e *echo.Echo, userController controller.UserController) {
	usersGroup := e.Group("users")

	usersGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	usersGroup.GET("", userController.GetAllUsers())
}
