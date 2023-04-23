package auth

import (
	"github.com/estradax/exater/internal/model/user"
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{
		"CreateSessionURL": ctx.OriginalURL(),
	})
}

type loginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func CreateSession(ctx *fiber.Ctx) error {
	// TODO: Validate data.
	p := new(loginForm)
	_ = ctx.BodyParser(p)

	u, found, _ := user.FindByEmail(p.Email)

	if !found {
		// TODO: Notify record is not exist.
		return ctx.Redirect(ctx.OriginalURL())
	}

	if u.Password != p.Password {
		// TODO: Credentials are not valid.
		return ctx.Redirect(ctx.OriginalURL())
	}

	sess, _ := session.SetUser(ctx, u)
	_ = sess.Save()

	accountURL, _ := ctx.GetRouteURL("account", fiber.Map{})
	redirectURI := ctx.Query("continue_uri", accountURL)

	return ctx.Redirect(redirectURI)
}
