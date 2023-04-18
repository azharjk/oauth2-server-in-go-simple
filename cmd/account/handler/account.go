package handler

import (
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
)

func Account(ctx *fiber.Ctx) error {
	user, _ := session.User(ctx)

	url, _ := ctx.GetRouteURL("logout", fiber.Map{})
	return ctx.Render("account", fiber.Map{
		"User":      user,
		"LogoutURL": url,
	})
}
