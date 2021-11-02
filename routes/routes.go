package routes

import (
	"auth-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/logged_in", controllers.LoggedIn)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/sensor_list", controllers.SensorList)
}
