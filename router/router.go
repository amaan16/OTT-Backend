package router

import "github.com/gofiber/fiber/v2"

// Define a function to set up the API endpoints
func SetupRoutes(app *fiber.App) {

	app.Get("/", helloWorld)
	// app.Get("/login", login)

	// userGroup := app.Group("/user")
	// userGroup.Get("/", getAllUsers)
	// userGroup.Get("/:id", getUserByID)
}

func helloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}
