package handler

import (
	"github.com/estradax/exater/internal/model"
	"github.com/estradax/exater/internal/model/user"
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx *fiber.Ctx) error {
	url, _ := ctx.GetRouteURL("registerUser", fiber.Map{})
	return ctx.Render("register", fiber.Map{
		"RegisterUserURL": url,
	})
}

type registerForm struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func RegisterUser(ctx *fiber.Ctx) error {
	// TODO: Validate data.
	p := new(registerForm)
	_ = ctx.BodyParser(p)

	_, found, _ := user.FindByEmail(p.Email)

	if found {
		// TODO: Notify record already exist.
		return ctx.RedirectToRoute("register", fiber.Map{})
	}

	u, _ := user.Create(model.User{
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
	})

	sess, _ := session.SetUser(ctx, u)
	_ = sess.Save()

	return ctx.RedirectToRoute("account", fiber.Map{})
}
