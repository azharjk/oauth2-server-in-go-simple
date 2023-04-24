package auth

import (
	"github.com/estradax/exater/internal/model"
	"github.com/estradax/exater/internal/model/user"
	"github.com/estradax/exater/internal/session"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx *fiber.Ctx) error {
	url, err := ctx.GetRouteURL("registerUser", fiber.Map{})
	if err != nil {
		return err
	}

	return ctx.Render("register", fiber.Map{
		"RegisterUserURL": url,
	})
}

type registerForm struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=8"`
}

func RegisterUser(ctx *fiber.Ctx) error {
	p := new(registerForm)
	if err := ctx.BodyParser(p); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return err
	}

	_, found, err := user.FindByEmail(p.Email)
	if err != nil {
		return err
	}

	if found {
		// TODO: Notify record already exist.
		return ctx.RedirectToRoute("register", fiber.Map{})
	}

	u, err := user.Create(model.User{
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
	})
	if err != nil {
		return err
	}

	sess, err := session.SetUser(ctx, u)
	if err != nil {
		return err
	}

	err = sess.Save()
	if err != nil {
		return err
	}

	return ctx.RedirectToRoute("account", fiber.Map{})
}
