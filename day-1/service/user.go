package service

import (
	"day-1/model"
	"day-1/repository"
	"errors"
)

type UserService interface {
	UserAvail(user model.User) (model.User, error)
	UserValid(user model.User) (model.User, error)
}

type userService struct {
	ur repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}


func (u *userService) UserAvail(user model.User) (model.User, error) {
	userRepo, err := u.ur.GetUserByUsername(user.Username)

	if err != nil {
		return model.User{}, err
	}

	if userRepo.Username != user.Username {
		return model.User{},errors.New("no users")
	}

	return userRepo, nil
}


func (u *userService) UserValid(user model.User) (model.User, error) {
	userRepo, err :=  u.ur.GetUserByUsername(user.Username)

	if err != nil {
		return model.User{}, err
	}

	if userRepo.Username != user.Username || userRepo.Password != user.Password {
		return model.User{}, errors.New("wrong username or password")
	}

	return userRepo, nil
}


