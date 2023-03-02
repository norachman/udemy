package main

import (
	"day-1/db"
	"day-1/handler"
	"log"

	// "github.com/NikSchaefer/go-fiber/database"
	// "github.com/NikSchaefer/go-fiber/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	
	err := db.Connection()
	if err!= nil {
		log.Fatal(err)
	}
	
	app := fiber.New()



	routes(app)
  	app.Listen(":8080")
}


func routes(app *fiber.App) {
	app.Get("/", handler.SayHello)
	app.Get("/id", handler.Returnjson)
	app.Get("/user", handler.GetUser)
}

