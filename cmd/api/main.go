package main

import (
	"awesomeProject1/internal/datadase"
	"awesomeProject1/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	datadase.Connection()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
