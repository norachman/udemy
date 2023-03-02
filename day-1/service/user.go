package service

import (
	"day-1/repository"
	"errors"
)

func GetUser() error {
	user := repository.GetAllUser()

	if len(user) == 0 {
		return errors.New("no user")
	}

	return nil
} 