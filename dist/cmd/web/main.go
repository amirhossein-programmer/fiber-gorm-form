package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/sixfwa/fiber-api/dist/cmd/web/database"
)

func main() {
	database.ConnectDb()
	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	engine := html.New("./ui/html/pages", ".html")
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "fiber",
		AppName:       "My App V1.0.1",
		Views:         engine,
	})
	app.Static("/static/", "./ui/static/", fiber.Static{
		ByteRange: true,
	})
	setupRoutes(app)
	log.Fatal(app.Listen(":4000"))
}
