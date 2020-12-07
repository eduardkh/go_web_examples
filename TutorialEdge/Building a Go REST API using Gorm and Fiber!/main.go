package main

import (
	"github.com/eduardkh/go_web_examples/tree/main/TutorialEdge/Building%20a%20Go%20REST%20API%20using%20Gorm%20and%20Fiber!/book"
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Put("/api/v1/book", book.UpdateBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
