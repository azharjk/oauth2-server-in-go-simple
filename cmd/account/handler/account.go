package handler

import (
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
)

func Account(ctx *fiber.Ctx) error {
	user, err := session.User(ctx)
	if err != nil {
		return err
	}

	url, err := ctx.GetRouteURL("logout", fiber.Map{})
	if err != nil {
		return err
	}

	return ctx.Render("account", fiber.Map{
		"User":      user,
		"LogoutURL": url,
	})
}
