package main

import ("github.com/gofiber/fiber/v2"
		"github.com/gofiber/fiber/v2/log"
		"github.com/gofiber/fiber/v2/middleware/cors"
)
func main() {
	// fmt.Println("Hello world!")
	app := fiber.New()

	app.Use(cors.new())
	app.Get("/api", func(c *fiber.Ctx) error {
			c.Status(200)
			return c.SendString("hello world")
	})
	log.Fatal(app.Listen("localhost:8000"))
}