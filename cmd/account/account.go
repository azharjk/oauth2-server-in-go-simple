package main

import (
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	_ = godotenv.Load()

	engine := html.New("./web/account", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	session.Setup()

	app.Use(logger.New())

	setupAccountRoute(app)
	setupAuthRoute(app)
	setupOAuth2Route(app)

	_ = app.Listen(os.Getenv("ACCOUNT_ADDR"))
}
