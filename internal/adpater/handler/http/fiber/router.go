package handler

import "github.com/gofiber/fiber/v2"

func Route() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "gukppap-backend",
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	urls := v1.Get("/urls")

	return app
}

func ErrorHandler() {

}
