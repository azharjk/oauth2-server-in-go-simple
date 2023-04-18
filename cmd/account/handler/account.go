package handler

import (
	"github.com/estradax/exater/internal/model"
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
)

func Account(ctx *fiber.Ctx) error {
	sess, _ := session.Get(ctx)
	user := sess.Get("user").(model.User)

	url, _ := ctx.GetRouteURL("logout", fiber.Map{})
	return ctx.Render("account", fiber.Map{
		"User":      user,
		"LogoutURL": url,
	})
}
