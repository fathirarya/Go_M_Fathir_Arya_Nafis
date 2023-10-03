package service

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(model.User) (model.User, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository}
}

func (service *UserServiceImpl) GetAllUsers() ([]model.User, error) {
	return service.UserRepository.FindAll()
}

func (service *UserServiceImpl) CreateUser(user model.User) (model.User, error) {
	return service.UserRepository.Save(user)
}
