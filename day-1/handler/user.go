package handler

import (
	"day-1/repository"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error{
	return c.JSON(repository.GetAllUser())
}