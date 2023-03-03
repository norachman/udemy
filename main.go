package main

import (
	"day-1/db"
	"day-1/handler"
	"day-1/repository"
	"day-1/service"
	"log"

	// "github.com/NikSchaefer/go-fiber/database"
	// "github.com/NikSchaefer/go-fiber/router"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPIHandler handler.UserAPI
}

func main() {

	err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	db := db.GetConnectionDB()

	api := NewInstance(db)

	app := fiber.New()

	routes(app, api)
	app.Listen(":8080")
}

func NewInstance(db *gorm.DB) APIHandler {
	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserAPI(userService)

	api := APIHandler{
		UserAPIHandler: userHandler,

	}

	return api
}

func routes(app *fiber.App, api APIHandler) {

	app.Post("/register", api.UserAPIHandler.Register )
	app.Post("/login", api.UserAPIHandler.Login)
}
