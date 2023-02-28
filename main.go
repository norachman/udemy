package main

import (
	"fmt"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()


  	app.Listen(":8080")
}

func hellohandler(c *fiber.Ctx) error {
	err := c.Send([]byte("Hello, World!"))

	return err
}

func routes(app *fiber.App) {
	app.Get("/", hellohandler)
}

