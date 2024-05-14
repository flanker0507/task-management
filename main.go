package main

import (
	"github.com/gofiber/fiber/v2"
	"go-todo-app/configs"
	"go-todo-app/routes"
)

func main() {

	configs.BootApp()
	app := fiber.New()
	//init route
	routes.InitRoute(app)

	app.Listen(":8000")

}
