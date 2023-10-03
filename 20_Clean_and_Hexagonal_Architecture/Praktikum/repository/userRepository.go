package repository

import (
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	Save(model.User) (model.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB}
}

func (repository *UserRepositoryImpl) FindAll() ([]model.User, error) {
	users := make([]model.User, 0)
	err := repository.DB.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (repository *UserRepositoryImpl) Save(user model.User) (model.User, error) {
	err := repository.DB.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
