package service

import (
	"day-1/model"
	"day-1/repository"
	"errors"

)

type UserService interface {
	Login(user *model.User) (model.User, error)
	Register(user *model.User) (model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userrepo repository.UserRepository) UserService {
	return &userService{userrepo}
}


func (u *userService) Register(user *model.User) (model.User, error) {
	
	userinRepo, err := u.userRepo.GetUserByUsername(user.Username)

	if err != nil {
		return model.User{}, err
	}

	if userinRepo.Username == user.Username {
		return model.User{}, errors.New("username is already exist")
	}

	err = u.userRepo.StoreUser(*user)
	if err != nil {
		return model.User{}, err
	}
	

	return *user, nil
}


func (u *userService) Login(user *model.User) (model.User, error) {
	userRepo, err :=  u.userRepo.GetUserByUsername(user.Username)

	if err != nil {
		return model.User{}, err
	}

	if userRepo.Username != user.Username || userRepo.Password != user.Password {
		return model.User{}, errors.New("wrong username or password")
	}


	return userRepo, nil
}




