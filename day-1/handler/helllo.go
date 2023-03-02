package handler

import (
	"github.com/gofiber/fiber/v2"
)

func SayHello(c *fiber.Ctx) error {
	err := c.Send([]byte("Hello, World!"))

	return err
}

func Returnjson(c *fiber.Ctx) error {
	err := c.JSON(fiber.Map{
		"name"		: "noer",
		"message" 	: "success",
	})

	return err
}