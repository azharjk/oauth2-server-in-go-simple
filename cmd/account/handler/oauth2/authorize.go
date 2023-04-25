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
	if err := ctx.QueryParser(p); err != nil {
		return err
	}

	c, found, err := client.FindByID(p.ClientID)
	if err != nil {
		return err
	}

	if !found {
		u, err := newErrorRedirectURI(c.RedirectURI, "invalid_request")
		if err != nil {
			return err
		}

		return ctx.Redirect(u)
	}

	if p.RedirectURI != c.RedirectURI {
		u, err := newErrorRedirectURI(c.RedirectURI, "invalid_request")
		if err != nil {
			return err
		}

		return ctx.Redirect(u)
	}

	u, err := newRedirectURI(c.RedirectURI)
	if err != nil {
		return err
	}

	return ctx.Redirect(u)
}
