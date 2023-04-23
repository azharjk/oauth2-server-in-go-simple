package oauth2

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type authorizeQuery struct {
	ResponseType string `query:"response_type"`
	ClientID     string `query:"client_id"`
}

func Authorize(ctx *fiber.Ctx) error {
	// TODO: Validate data.
	p := new(authorizeQuery)
	_ = ctx.QueryParser(p)

	u, _ := url.Parse("http://localhost:9080/oauth2/callback")
	q := u.Query()
	q.Set("code", "code")
	u.RawQuery = q.Encode()

	return ctx.Redirect(u.String())
}
