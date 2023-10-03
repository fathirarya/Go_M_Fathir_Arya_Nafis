package controller

import (
	"belajar-go-echo/helper"
	"belajar-go-echo/model"
	"belajar-go-echo/service"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	CreateUser() echo.HandlerFunc
	GetAllUsers() echo.HandlerFunc
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{userService}
}

func (controller *UserControllerImpl) CreateUser() echo.HandlerFunc {
	return func(echoContext echo.Context) error {
		user := model.User{}
		err := echoContext.Bind(&user)
		if err != nil {
			return echoContext.JSON(500, map[string]interface{}{
				"message": err.Error(),
			})
		}
		newUser, err := controller.UserService.CreateUser(user)
		if err != nil {
			return echoContext.JSON(500, map[string]interface{}{
				"message": err.Error(),
			})
		}
		token, err := helper.GenerateToken(newUser.ID, newUser.Email)
		if err != nil {
			return echoContext.JSON(500, map[string]interface{}{
				"message": err.Error(),
			})

		}
		response := struct {
			User  model.User `json:"user"`
			Token string     `json:"token"`
		}{
			newUser,
			token,
		}
		return echoContext.JSON(200, response)
	}
}

func (controller *UserControllerImpl) GetAllUsers() echo.HandlerFunc {
	return func(echoContext echo.Context) error {
		users, err := controller.UserService.GetAllUsers()
		if err != nil {
			return echoContext.JSON(500, map[string]interface{}{
				"message": err.Error(),
			})
		}
		return echoContext.JSON(200, users)
	}
}
