package handler

import (
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func Authorized(ctx *fiber.Ctx) error {
	b, err := session.IsAuthorized(ctx)
	if err != nil {
		return err
	}

	queries := map[string]string{
		"continue_uri": url.QueryEscape(ctx.BaseURL() + ctx.OriginalURL()),
	}

	if !b {
		return ctx.RedirectToRoute("login", fiber.Map{
			"queries": queries,
		})
	}

	return ctx.Next()
}
