package auth

import (
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
)

func Logout(ctx *fiber.Ctx) error {
	sess, err := session.Get(ctx)
	if err != nil {
		return err
	}

	err = sess.Regenerate()
	if err != nil {
		return err
	}

	return ctx.RedirectToRoute("login", fiber.Map{})
}
