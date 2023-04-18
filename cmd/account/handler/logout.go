package handler

import (
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
)

func Logout(ctx *fiber.Ctx) error {
	sess, _ := session.Get(ctx)
	_ = sess.Regenerate()

	return ctx.RedirectToRoute("login", fiber.Map{})
}
