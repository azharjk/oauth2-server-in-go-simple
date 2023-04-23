package oauth2

import (
	"github.com/gofiber/fiber/v2"
	_oauth2 "golang.org/x/oauth2"
	"time"
)

type tokenBody struct {
	GrantType string `form:"grant_type"`
	Code      string `form:"code"`
}

func Token(ctx *fiber.Ctx) error {
	p := new(tokenBody)
	_ = ctx.BodyParser(p)

	tok := _oauth2.Token{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		Expiry:       time.Now().Add(3600),
		TokenType:    "bearer",
	}

	return ctx.JSON(tok)
}
