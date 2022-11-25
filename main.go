package main

import (
	"go-checkout/src/database"
	"go-checkout/src/routes"
	"go-checkout/src/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()
    app := fiber.New()
	services.Setup()
	// events.SetupProducer()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	

    log.Fatal(app.Listen(":8000"))
}