package session

import (
	"github.com/estradax/exater/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

func Setup() {
	store.RegisterType(model.User{})
}

func Get(ctx *fiber.Ctx) (*session.Session, error) {
	return store.Get(ctx)
}
