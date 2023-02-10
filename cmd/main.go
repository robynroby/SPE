package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robynroby/SPE/database"
)

func main() {
	// conncet to database
	database.ConnectDb()

	app := fiber.New()

	// setup routes
	setupRoutes(app)

	app.Listen(":3000")
}
