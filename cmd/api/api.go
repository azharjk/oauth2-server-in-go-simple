package main

import (
	"fmt"
	"github.com/estradax/exater/cmd/api/handler/oauth2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	_ = godotenv.Load()

	app := fiber.New()

	app.Use(logger.New())

	basic := basicauth.New(basicauth.Config{
		Authorizer: func(id, secret string) bool {
			return true
		},
	})

	app.Get("/userinfo", func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Get(fiber.HeaderAuthorization))
		return ctx.JSON(fiber.Map{
			"id":   "name",
			"name": "user",
		})
	})

	app.Post("/oauth2/token", basic, oauth2.Token)

	_ = app.Listen(os.Getenv("API_ADDR"))
}
