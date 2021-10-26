package main

import (
	"auth-go/database"

	"auth-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	//fiber
	app := fiber.New()

	//para aceder através do browser(sem ser só pela porta 8000)
	app.Use(cors.New(cors.Config{

		AllowCredentials: true, //para obter o cookie
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
