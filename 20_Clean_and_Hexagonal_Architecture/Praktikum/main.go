package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"belajar-go-echo/routes"
	"belajar-go-echo/service"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	db, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.User{})

	userRepository := repository.NewUserRepository(db)
	userServcice := service.NewUserService(userRepository)
	userController := controller.NewUserController(userServcice)

	routes.RoutesWithoutController(app, userController)
	routes.RoutesWithMiddleware(app, userController)

	app.Logger.Fatal(app.Start(":8080"))
}
