package main

import (
	"github.com/estradax/exater/cmd/account/handler"
	"github.com/estradax/exater/internal/model"
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func main() {
	model.Setup()

	engine := html.New("./web/account", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	session.Setup()

	app.Use(logger.New())

	app.Get("/", handler.Authorized, handler.Account).Name("account")

	setupAuthRoute(app)

	_ = app.Listen(":8080")
}
