package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func SetResponse(m string, data interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["messages"] = m
	response["data"] = data

	return response

}

// GET ALL USER
func GetUsersControllers(c echo.Context) error {
	return c.JSON(http.StatusOK, SetResponse("success get all users", users))
}

// GET USER BY ID
func GetUserController(c echo.Context) error {
	//BAD REQUEST
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse("invalid user id", nil))
	}

	var getUserById *User
	for i, user := range users {
		if user.Id == id {
			getUserById = &users[i]
			break
		}
	}

	//USER NOT FOUND
	if getUserById == nil {
		return c.JSON(http.StatusNotFound, SetResponse("user not found", nil))
	}

	//SUCCESS GET USER BY ID
	return c.JSON(http.StatusOK, SetResponse("successgrt user by id", getUserById))

}

// //DELETE USER BY ID
func DeleteUserController(c echo.Context) error {
	//BAD REQUEST
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse("invalid delete user", nil))

	}

	var deletedUser *User
	for i, user := range users {
		if user.Id == id {
			deletedUser = &users[i]
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	// USER NOT FOUND
	if deletedUser == nil {
		return c.JSON(http.StatusNotFound, SetResponse("user not found", nil))
	}

	// SUCCESS DELETE USER
	return c.JSON(http.StatusOK, SetResponse("success delete user", deletedUser))
}

// UPDATE USER BY ID
func UpdateUserController(c echo.Context) error {
	// BINDING DATA
	user := User{}
	c.Bind(&user)

	// BAD REQUEST
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse("invalid update user", nil))
	}

	var updateUser *User
	for i, user := range users {
		if user.Id == id {
			updateUser = &users[i]
			updateUser.Name = user.Name
			updateUser.Email = user.Email
			updateUser.Password = user.Password
			break
		}
	}

	// USER NOT FOUND
	if updateUser == nil {
		return c.JSON(http.StatusNotFound, SetResponse("user not found", nil))
	}

	// SUCCESS UPDATE USER
	return c.JSON(http.StatusOK, SetResponse("success update user", updateUser))

}

// CREATE NEW USER
func CreateUserController(c echo.Context) error {
	// BINDING DATA
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, SetResponse("success create user", users))
}
