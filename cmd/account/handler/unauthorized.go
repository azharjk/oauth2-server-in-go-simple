package handler

import (
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
)

func Unauthorized(ctx *fiber.Ctx) error {
	sess, _ := session.Get(ctx)
	user := sess.Get("user")

	if user != nil {
		return ctx.RedirectToRoute("account", fiber.Map{})
	}

	return ctx.Next()
}
