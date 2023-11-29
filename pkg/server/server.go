package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Start() {
	app := fiber.New(fiber.Config{
		AppName: "put-call-ratio",
	})
	app.Use(favicon.New())
	app.Use(cors.New())
	app.Use(recover.New())

	app.Static("/", "./public")

	loadRoutes(app)

	log.Printf("Starting server on port 80..")
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}

	defer func() {
		log.Printf("Shutting down app..")
		app.Shutdown()
	}()
}
