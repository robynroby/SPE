package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robynroby/SPE/handlers"
)

func setupRoutes(app *fiber.App){
	app.Get("/facts",handlers.ListFacts)
	app.Post("/addfact",handlers.CreateFact)
}