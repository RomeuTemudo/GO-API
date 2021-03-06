package routes

import (
	"auth-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/add_user", controllers.AddUser)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/logged_in", controllers.LoggedIn)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/sensor_list", controllers.SensorList)
	app.Post("/api/add_sensor", controllers.AddSensor)
	app.Post("/api/update_sensor", controllers.UpdateSensor)
	app.Get("/api/categories", controllers.GetCategories)
	app.Post("/api/delete_sensor", controllers.DeleteSensor)
	app.Get("/api/search_sensor", controllers.SensorSearch)
	app.Post("/api/add_sensor_data", controllers.AddSensorData)
	app.Get("/api/get_sensor_data", controllers.GetSensorData)
	app.Get("/api/get_roles", controllers.GetRoles)
	app.Get("/api/get_users", controllers.GetUsers)
	app.Post("/api/delete_user", controllers.DeleteUser)
}
