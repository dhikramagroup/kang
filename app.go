package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./views/assets")
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Jasa tukang bangun rumah!",
		})
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("about", fiber.Map{
			"Title": "Tentang kang bangunan!",
		})
	})

	app.Get("/contact-us", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("contact", fiber.Map{
			"Title": "Kontak kang bangunan!",
		})
	})

	app.Get("/gallery", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("gallery", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/blog", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("blog", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	log.Fatal(app.Listen(":8080"))
}
