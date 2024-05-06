package handlers

import (
	"github.com/Rabbit937/divrhino-trivia/database"
	"github.com/Rabbit937/divrhino-trivia/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Div Rhino Trivia App!1111111")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}