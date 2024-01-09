package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/guckppap/gukppap-backend/internal/adpater/handler/http/fiber/dto"
	"github.com/guckppap/gukppap-backend/pkg/apperrors"
)

type Router struct {
	app *fiber.App
}

func (r *Router) Listen(port string) error {
	return r.app.Listen(port)
}

func NewRoute(shortcutHandler *ShortcutHandler) *Router {
	app := fiber.New(fiber.Config{
		AppName:      "gukppap-backend",
		ErrorHandler: ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/:shortcut", shortcutHandler.Redirect)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	urls := v1.Group("/urls")

	urls.Get("/shortcut", shortcutHandler.Get)
	urls.Post("/shortcut", shortcutHandler.Save)
	urls.Patch("/shortcut", shortcutHandler.Update)
	urls.Delete("/shortcut", shortcutHandler.Delete)

	return &Router{app: app}
}

func ErrorHandler(c *fiber.Ctx, err error) error {

	if apperrors.IsEqual(apperrors.BadReqeustError, err) {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.General{
			Message: err.Error(),
		})
	}

	if apperrors.IsEqual(apperrors.InternalServerError, err) {
		return c.Status(fiber.StatusInternalServerError).JSON(&dto.General{
			Message: err.Error(),
		})
	}

	if apperrors.IsEqual(apperrors.NotFoundError, err) {
		return c.Status(fiber.StatusNotFound).JSON(&dto.General{
			Message: err.Error(),
		})
	}

	if apperrors.IsEqual(apperrors.RequestParsingError, err) {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.General{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusNotFound).JSON(dto.General{
		Message: err.Error(),
	})
}
