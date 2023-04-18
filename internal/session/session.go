package session

import (
	"github.com/estradax/exater/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetUser(ctx *fiber.Ctx, user model.User) (*session.Session, error) {
	sess, err := Get(ctx)
	sess.Set("user", user)
	return sess, err
}