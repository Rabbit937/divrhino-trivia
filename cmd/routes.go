package main

import (
	"github.com/Rabbit937/divrhino-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	
	app.Post("/fact", handlers.CreateFact)
}
