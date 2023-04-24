package auth

import (
	"github.com/estradax/exater/internal/model/user"
	"github.com/estradax/exater/internal/session"
	"github.com/estradax/exater/internal/validate"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{
		"CreateSessionURL": ctx.OriginalURL(),
	})
}

type loginForm struct {
	Email    string `form:"email" validate:"required"`
	Password string `form:"password" validate:"required"`
}

func CreateSession(ctx *fiber.Ctx) error {
	p := new(loginForm)
	if err := ctx.BodyParser(p); err != nil {
		return err
	}

	if err := validate.Struct(p); err != nil {
		return err
	}

	u, found, err := user.FindByEmail(p.Email)
	if err != nil {
		return err
	}

	if !found {
		sess, err := session.SetError(ctx, "record is not found")
		if err != nil {
			return err
		}

		if err := sess.Save(); err != nil {
			return err
		}

		return ctx.Redirect(ctx.OriginalURL())
	}

	if u.Password != p.Password {
		sess, err := session.SetError(ctx, "credentials is not valid")
		if err != nil {
			return err
		}

		if err := sess.Save(); err != nil {
			return err
		}

		return ctx.Redirect(ctx.OriginalURL())
	}

	sess, err := session.SetUser(ctx, u)
	if err != nil {
		return err
	}

	err = sess.Save()
	if err != nil {
		return err
	}

	accountURL, err := ctx.GetRouteURL("account", fiber.Map{})
	if err != nil {
		return err
	}

	redirectURI := ctx.Query("continue_uri", accountURL)

	return ctx.Redirect(redirectURI)
}
