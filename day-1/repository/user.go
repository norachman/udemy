package repository

import (
	"day-1/model"
	"encoding/base64"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]model.User, error)
	GetUserByUsername(username string) (model.User, error)
	StoreUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) GetAllUser() ([]model.User, error) {
	var user []model.User

	resp := u.db.Table("users").Find(&user)
	if resp.Error != nil {
		return []model.User{}, resp.Error
	}

	return user, nil
}

func (u *userRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User

	resp := u.db.Table("users").Where("username = ?", username).Find(&user)
	if resp.Error != nil {
		return model.User{}, resp.Error
	}

	pwd, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil {
		return model.User{}, nil
	}
	user.Password = string(pwd)

	return user, nil
}

func (u *userRepository) StoreUser(user *model.User) error {
	pwd := base64.StdEncoding.EncodeToString([]byte(user.Password))
	user.Password = pwd

	resp := u.db.Table("users").Save(&user)
	if resp.Error != nil {
		return resp.Error
	}

	return nil
}

func (u *userRepository) UpdateUser(user *model.User) error {
	resp := u.db.Table("users").Updates(user)
	if resp.Error != nil {
		return resp.Error
	}

	return nil
}

func (u *userRepository) DeleteUser(id int) error {
	resp := u.db.Table("users").Where("id = ?", id).Delete(&model.User{})

	if resp.Error != nil {
		return resp.Error
	}

	return nil
}
