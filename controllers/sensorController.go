package controllers

import (
	"auth-go/database"
	"auth-go/models"

	"github.com/gofiber/fiber/v2"
)

func SensorList(c *fiber.Ctx) error {

	var sensors []models.Sensor

	database.DB.Find(&sensors)

	return c.JSON(sensors)

}

// GetAllProducts from db
/*func GetAllProducts(c *fiber.Ctx) {
    // query product table in the database
    rows, err := database.DB.Query("SELECT name, description, category, amount FROM products order by name")
    if err != nil {
        c.Status(500).JSON(&fiber.Map{
            "success": false,
            "error": err,
          })
        return
    }
    defer rows.Close()
    result := model.Products{}
    for rows.Next() {
        product := model.Product{}
        err := rows.Scan(&product.Name, &product.Description, &product.Category, &product.Amount)
        // Exit if we get an error
        if err != nil {
            c.Status(500).JSON(&fiber.Map{
                "success": false,
                "error": err,
              })
            return
        }
        // Append Product to Products
        result.Products = append(result.Products, product)
    }
    // Return Products in JSON format
    if err := c.JSON(&fiber.Map{
        "success": true,
        "product":  result,
        "message": "All product returned successfully",
      }); err != nil {
        c.Status(500).JSON(&fiber.Map{
            "success": false,
            "message": err,
          })
        return
    }
}*/
