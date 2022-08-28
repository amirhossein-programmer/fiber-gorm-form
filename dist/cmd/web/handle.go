package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/sixfwa/fiber-api/dist/cmd/web/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.Render("form", fiber.Map{
		"title": "goform",
	})
}
func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	// app.Get("/static", fileServer)
}
