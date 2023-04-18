package handler

import (
	"github.com/estradax/exater/internal/model"
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

	user := new(model.User)
	result := model.DB.Limit(1).Find(user, model.User{Email: p.Email})
	_ = result.Error

	if result.RowsAffected != 1 {
		// TODO: Notify record is not exist.
		return ctx.RedirectToRoute("login", fiber.Map{})
	}

	if user.Password != p.Password {
		// TODO: Credentials are not valid.
		return ctx.RedirectToRoute("login", fiber.Map{})
	}

	sess, _ := session.Get(ctx)
	sess.Set("user", user)
	_ = sess.Save()

	return ctx.RedirectToRoute("account", fiber.Map{})
}
