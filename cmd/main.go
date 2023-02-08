package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robynroby/SPE/database"
)

func main() {
	database.Connect()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
