package main

import (
	"OTT/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Set up the API routes by calling the setupRoutes function from router.go
	router.SetupRoutes(app)

	// Start the server on port 3000
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
