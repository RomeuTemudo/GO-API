package controllers

import (
	"auth-go/database"
	"auth-go/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SensorList(c *fiber.Ctx) error {

	var sensors []models.Sensor

	param_id := c.Query("id")

	if param_id == "" {

		database.DB.Find(&sensors)
	} else {

		database.DB.First(&sensors, param_id)

	}

	return c.JSON(sensors)

}

func SensorSearch(c *fiber.Ctx) error {

	var sensors []models.Sensor

	param_name := c.Query("name")

	database.DB.Where("sensor_name LIKE ?", "%"+param_name+"%").Find(&sensors)

	return c.JSON(sensors)

}

func AddSensor(c *fiber.Ctx) error {

	sensor := new(models.Sensor)

	if err := c.BodyParser(sensor); err != nil {
		return err
	}

	database.DB.Create(&sensor)

	table_name := "sensor_data_" + strconv.Itoa(sensor.SensorID)

	database.DB.Table(table_name).AutoMigrate(&models.SensorData{})

	return c.JSON(&sensor)

}

func UpdateSensor(c *fiber.Ctx) error {

	sensor := new(models.Sensor)

	if err := c.BodyParser(sensor); err != nil {
		return err
	}

	//database.DB.Model(&sensor).Where("sensor_id=?", sensor.SensorID).Select("sensor_name", "sensor_description", "category_id").Updates(models.Sensor{SensorName: sensor.SensorName, SensorDescription: sensor.SensorDescription, CategoryID: sensor.CategoryID})

	database.DB.Where("sensor_id=?", sensor.SensorID).Updates(models.Sensor{SensorName: sensor.SensorName, SensorDescription: sensor.SensorDescription, CategoryID: sensor.CategoryID})

	return c.JSON(fiber.Map{
		"Status":  "Success",
		"Message": "Boa burro , criaste um sensor!",
	})

}

func DeleteSensor(c *fiber.Ctx) error {

	sensor := new(models.Sensor)

	if err := c.BodyParser(sensor); err != nil {
		return err
	}

	//database.DB.Model(&sensor).Where("sensor_id=?", sensor.SensorID).Delete(&sensor)

	//database.DB.Delete(models.Sensor{SensorID: sensor.SensorID}, sensor.SensorID)

	database.DB.Delete(&sensor, sensor.SensorID)

	return c.JSON(fiber.Map{
		"Status":  "Success",
		"Message": "Boa burro , apagaste um sensor!",
	})

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
