package handler

import (
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
)

func Error(ctx *fiber.Ctx) error {
	sess, err := session.Get(ctx)
	if err != nil {
		return err
	}

	e := sess.Get("error")

	if err := ctx.Bind(fiber.Map{
		"error": e,
	}); err != nil {
		return err
	}

	return ctx.Next()
}
