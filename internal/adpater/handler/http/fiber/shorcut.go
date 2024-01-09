package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guckppap/gukppap-backend/internal/adpater/handler/http/fiber/dto"
	"github.com/guckppap/gukppap-backend/internal/core/ports"
	"github.com/guckppap/gukppap-backend/pkg/apperrors"
	"github.com/guckppap/gukppap-backend/pkg/verifier"
)

type ShortcutHandler struct {
	shortcutService ports.ShortcutService
}

func NewShortcutHandler(shortcutService ports.ShortcutService) *ShortcutHandler {
	return &ShortcutHandler{shortcutService: shortcutService}
}

// [GET] /api/v1/urls/shortcut
func (sh *ShortcutHandler) Get(c *fiber.Ctx) error {
	// body parsing
	body := new(dto.GetSaveDelete)
	if err := c.BodyParser(body); err != nil {
		return apperrors.New(apperrors.BadReqeustError, "Unable to parse."+err.Error())
	}

	// 유효성 검사
	if err := verifier.Run(body); err != nil {
		return err
	}

	// 비지니스 로직 호출
	shorcut, err := sh.shortcutService.GetByUrl(c.Context(), body.Url)
	if err != nil {
		return err
	}

	// 정상적으로 res
	return c.Status(fiber.StatusOK).JSON(&dto.General{
		Data: shorcut,
	})
}

// [POST] /api/v1/urls/shortcut
func (sh *ShortcutHandler) Save(c *fiber.Ctx) error {
	// body parsing
	body := new(dto.GetSaveDelete)
	if err := c.BodyParser(body); err != nil {
		return apperrors.New(apperrors.BadReqeustError, "Unable to parse."+err.Error())
	}

	// 유효성 검사
	if err := verifier.Run(body); err != nil {
		return err
	}

	// 비지니스 로직 호출
	shorcut, err := sh.shortcutService.Save(c.Context(), body.Url)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(&dto.General{
		Data: shorcut,
	})
}

// [PATCH] /api/v1/urls/shortcut
func (sh *ShortcutHandler) Update(c *fiber.Ctx) error {
	// body parsing
	body := new(dto.Update)
	if err := c.BodyParser(body); err != nil {
		return apperrors.New(apperrors.BadReqeustError, "Unable to parse."+err.Error())
	}

	// 유효성 검사
	if err := verifier.Run(body); err != nil {
		return err
	}

	// 비지니스 로직 호출
	if err := sh.shortcutService.Update(c.Context(), body.Shortcut, body.Want); err != nil {
		return err
	}

	// 응답
	return c.SendStatus(fiber.StatusOK)

}

// [DELETE] /api/v1/urls/shortcut
func (sh *ShortcutHandler) Delete(c *fiber.Ctx) error {
	// body parsing
	body := new(dto.GetSaveDelete)
	if err := c.BodyParser(body); err != nil {
		return apperrors.New(apperrors.BadReqeustError, "Unable to parse."+err.Error())
	}

	// 유효성 검사
	if err := verifier.Run(body); err != nil {
		return err
	}

	// 비지니스 로직 호출
	if err := sh.shortcutService.Delete(c.Context(), body.Url); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
