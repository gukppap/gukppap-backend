package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type router struct {
	router *fiber.App
}

func NewRouter(urlH *urlHandler) *router {

	app := fiber.New()

	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			urls := v1.Group("/urls")
			{
				urls.Get("/shortcut", urlH.getURLByOriginalURL)
				urls.Get("original", urlH.getURLByShortcutURL)
				urls.Post("/shortcut", urlH.createURL)
				urls.Patch("/history", urlH.UpdateURL)
				urls.Delete("/history", urlH.DeleteURL)
			}

		}
	}

	return &router{app}
}

func (r *router) Run(port int) error {
	return r.router.Listen(fmt.Sprintf(":%d", port))
}
