package handler

import (
	"github.com/estradax/exater/internal/model/user"
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	url, _ := ctx.GetRouteURL("createSession", fiber.Map{})
	return ctx.Render("login", fiber.Map{
		"CreateSessionURL": url,
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
		return ctx.RedirectToRoute("login", fiber.Map{})
	}

	if u.Password != p.Password {
		// TODO: Credentials are not valid.
		return ctx.RedirectToRoute("login", fiber.Map{})
	}

	sess, _ := session.SetUser(ctx, u)
	_ = sess.Save()

	return ctx.RedirectToRoute("account", fiber.Map{})
}
