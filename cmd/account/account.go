package main

import (
	"github.com/estradax/exater/cmd/account/handler"
	"github.com/estradax/exater/internal/model"
	"github.com/estradax/exater/internal/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := model.Connect(); err != nil {
		log.Fatal(err)
	}

	engine := html.New("./web/account", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	session.Setup()

	app.Use(logger.New())
	app.Use(handler.Error)

	setupAccountRoute(app)
	setupAuthRoute(app)
	setupOAuth2Route(app)

	if err := app.Listen(os.Getenv("ACCOUNT_ADDR")); err != nil {
		log.Fatal(err)
	}
}
