package handler

import (
	"github.com/estradax/exater/internal/model"
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

	result := model.DB.Limit(1).Find(&model.User{}, model.User{Email: p.Email})
	_ = result.Error

	if result.RowsAffected != 0 {
		// TODO: Notify record already exist.
		return ctx.RedirectToRoute("register", fiber.Map{})
	}

	user := model.User{Name: p.Name, Email: p.Email, Password: p.Password}
	result = model.DB.Create(&user)
	_ = result.Error

	sess, _ := session.Get(ctx)
	sess.Set("user", user)
	_ = sess.Save()

	return ctx.RedirectToRoute("account", fiber.Map{})
}
