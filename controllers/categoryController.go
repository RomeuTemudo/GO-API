package controllers

import (
	"auth-go/database"
	"auth-go/models"

	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {

	var categories []models.Category

	database.DB.Find(&categories)

	return c.JSON(categories)

}
