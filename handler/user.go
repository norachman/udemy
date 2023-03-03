package handler

import (
	"day-1/model"
	"day-1/service"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserAPI interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userservice service.UserService) UserAPI {
	return &userAPI{userservice}
}

func (u *userAPI) Login(c *fiber.Ctx) error {
	var user model.UserLogin
	
	err := json.Unmarshal(c.Request().Body(), &user)
	
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(model.UserResponse{
			Message: "invalid decode json",
		})
		return err
	}

	if user.Username == "" || user.Password == ""{
		c.Status(fiber.StatusBadRequest).JSON(model.UserResponse{
			Message: "login data is empty",
		})
		return err
	}

	_, err = u.userService.Login(&model.User{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(model.UserResponse{
			Message: err.Error(),
		})
		return err
	}

	c.Cookie(
		&fiber.Cookie{
			Name: "login_session",
			Path: "/",
			Expires: time.Now().Add(5 * time.Hour),
		},
	)

	c.Status(200).JSON(map[string]interface{}{
		"username"	: user.Username,
		"message"	: "login success",
	})
	return nil
}

func (u *userAPI) Register(c *fiber.Ctx) error {
	var user model.UserRegister

	err := json.Unmarshal(c.Request().Body(), &user)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(model.UserResponse{
			Message: "invalid decode json",
		})
		return err
	}

	if user.Username == "" || user.Password == "" {
		c.Status(fiber.StatusBadRequest).JSON(model.UserResponse{
			Message: "register data is empty",
		})
		return err
	}

	_, err = u.userService.Register(&model.User{
		Username: user.Username,
		Fullname: user.Fullname,
		Password: user.Password,
	})
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(model.UserResponse{
			Message: err.Error(),
		})
		return err
	}

	c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"username"	: user.Username,
		"message"	: "register success",
	})
	return nil



}


