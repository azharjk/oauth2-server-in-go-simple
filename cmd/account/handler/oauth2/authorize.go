package oauth2

import (
	"github.com/estradax/exater/internal/model/client"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type authorizeQuery struct {
	ResponseType string `query:"response_type"`
	ClientID     string `query:"client_id"`
	RedirectURI  string `query:"redirect_uri"`
}

func newErrorRedirectURI(s, s1 string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("error", s1)
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func newRedirectURI(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("code", "code")
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func Authorize(ctx *fiber.Ctx) error {
	// TODO: Validate data.
	p := new(authorizeQuery)
	_ = ctx.QueryParser(p)

	c, _, _ := client.FindByID(p.ClientID)

	if p.RedirectURI != c.RedirectURI {
		u, _ := newErrorRedirectURI(c.RedirectURI, "invalid_request")
		return ctx.Redirect(u)
	}

	u, _ := newRedirectURI(c.RedirectURI)

	return ctx.Redirect(u)
}
