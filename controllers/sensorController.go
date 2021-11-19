package controllers

import (
	"auth-go/database"
	"auth-go/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func SensorList(c *fiber.Ctx) error {

	var sensors []models.Sensor

	param_id := c.Query("id")
	param_order := c.Query("sort")

	//jesus christ :D
	if param_order == "true" {
		param_order = "sensor_name asc"

	}
	if param_order == "false" {

		param_order = "sensor_name desc"
	}

	if param_order == "undefined" {

		param_order = ""
	}

	if param_id == "" {

		database.DB.Order(param_order).Find(&sensors)
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

	//create dynamic table for sensor data(sensor_data + id)
	table_name := "sensor_data_" + strconv.Itoa(sensor.SensorID)

	database.DB.Table(table_name).AutoMigrate(&models.SensorValue{})

	return c.JSON(&sensor)

}

func AddSensorData(c *fiber.Ctx) error {

	var table_name string

	sensorData := new(models.SensorData)

	sensor := new(models.Sensor)

	if err := c.BodyParser(sensorData); err != nil {
		return err
	}

	table_name = "sensor_data_" + strconv.Itoa(sensorData.ID)

	//on create doesn't have an ID
	if sensorData.ID == 0 {

		database.DB.Table("sensors").Last(&sensor)

		lastSensorId := sensor.SensorID

		table_name = "sensor_data_" + strconv.Itoa(lastSensorId)

	}

	database.DB.Table(table_name).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&sensorData.Data)

	return c.JSON(fiber.Map{
		"Status":  "Success",
		"Message": "Boa burro , criaste um sensor!",
	})
}

func GetSensorData(c *fiber.Ctx) error {

	var sensorData []models.SensorValue

	param_id := c.Query("id")

	var table_name = "sensor_data_" + param_id

	database.DB.Table(table_name).Find(&sensorData)

	return c.JSON(sensorData)

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
