package handler

import (
	"gukppap-backend/internal/adapter/handler/http/fiber/dto"
	"gukppap-backend/internal/adapter/repository/mysql/ent"
	"gukppap-backend/internal/core/domain"
	"gukppap-backend/internal/core/service"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type urlHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *urlHandler {
	return &urlHandler{service}
}

func (uh *urlHandler) getURLByOriginalURL(c *fiber.Ctx) error {

	body := new(dto.URLDTO)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusBadRequest,
			Message: "unable http body",
		})
	}

	// 유효성 검사
	if _, err := url.ParseRequestURI(body.OriginalURL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusBadRequest,
			Message: "Not valid url",
		})
	}

	res, err := uh.service.GetURLByOrignalURL(c.Context(), body.OriginalURL)

	if _, ok := err.(*ent.NotFoundError); ok {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusBadRequest,
			Message: "not found url",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusInternalServerError,
			Message: "server error: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralDTO{
		Status:  fiber.StatusOK,
		Message: "succesed get",
		Data: dto.URLDTO{
			ShortcutURL: res.ShortcutURL,
		},
	})

}

func (uh *urlHandler) createURL(c *fiber.Ctx) error {

	body := new(dto.URLDTO)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusBadRequest,
			Message: "unable http body",
		})
	}

	// 유효성 검사
	if _, err := url.ParseRequestURI(body.OriginalURL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusBadRequest,
			Message: "Not valid url",
		})
	}

	res, err := uh.service.CreateURL(c.Context(), &domain.URL{OriginalURL: body.OriginalURL})

	// 중복인가?
	if _, ok := err.(*ent.ConstraintError); ok {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusBadRequest,
			Message: "Duplicate originalURL",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusInternalServerError,
			Message: "server error: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&dto.GeneralDTO{
		Status:  fiber.StatusCreated,
		Message: "successed create shorcut url",
		Data: &dto.URLDTO{
			ShortcutURL: res.ShortcutURL,
		},
	})

}

func (uh *urlHandler) UpdateURL(c *fiber.Ctx) error {

	body := new(dto.URLDTO)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusBadRequest,
			Message: "unable http body",
		})
	}

	// 유효성 검사
	if _, err := url.ParseRequestURI(body.OriginalURL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusBadRequest,
			Message: "Not valid url",
		})
	}

	if _, err := uh.service.UpdateURL(c.Context(), body.ShortcutURL, &domain.URL{
		OriginalURL: body.OriginalURL,
		ShortcutURL: body.ShortcutURL,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusInternalServerError,
			Message: "server error: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralDTO{
		Status:  fiber.StatusOK,
		Message: "successed update",
	})

}

func (uh *urlHandler) DeleteURL(c *fiber.Ctx) error {

	body := new(dto.URLDTO)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusBadRequest,
			Message: "unable http body",
		})
	}

	// 유효성 검사
	if _, err := url.ParseRequestURI(body.OriginalURL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusBadRequest,
			Message: "Not valid url",
		})
	}

	if err := uh.service.DeleteByOriginalURL(c.Context(), body.OriginalURL); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&dto.GeneralDTO{
			Status:  fiber.StatusInternalServerError,
			Message: "server error: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralDTO{
		Status:  fiber.StatusOK,
		Message: "successed delete",
	})

}
