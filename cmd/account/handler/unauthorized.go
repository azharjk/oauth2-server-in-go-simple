package handler

import (
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
)

func Unauthorized(ctx *fiber.Ctx) error {
	b, err := session.IsAuthorized(ctx)
	if err != nil {
		return err
	}

	if b {
		return ctx.RedirectToRoute("account", fiber.Map{})
	}

	return ctx.Next()
}
