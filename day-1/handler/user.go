package handler

import (
	"day-1/service"
	"github.com/gofiber/fiber/v2"
)

type UserAPI interface {

}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userservice service.UserService) UserAPI {
	return &userAPI{userservice}
}

func (u *userAPI) GetAllUser(c *fiber.Ctx) error{
	return c.JSON(u.userService.UserAvail())
}


