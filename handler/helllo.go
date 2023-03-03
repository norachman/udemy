package handler

import (
	"github.com/gofiber/fiber/v2"
)

type HelloAPI interface {
	SayHello(c *fiber.Ctx) error
	Returnjson(c *fiber.Ctx) error
}

type helloAPI struct{
}

func NewHelloHandler() helloAPI {
	return helloAPI{}
}

func (h *helloAPI) SayHello(c *fiber.Ctx) error {
	err := c.Send([]byte("Hello, World!"))

	return err
}

func (h *helloAPI) Returnjson(c *fiber.Ctx) error {
	err := c.JSON(fiber.Map{
		"name"		: "noer",
		"message" 	: "success",
	})

	return err
}