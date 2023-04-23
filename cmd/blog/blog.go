package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	_ = godotenv.Load()

	engine := html.New("./web/blog", ".html")

	conf := &oauth2.Config{
		ClientID:     "5",
		ClientSecret: "YOUR_CLIENT_SECRET",
		RedirectURL:  "http://" + os.Getenv("BLOG_ADDR") + "/oauth2/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:   "http://" + os.Getenv("ACCOUNT_ADDR") + "/oauth2/authorize",
			TokenURL:  "http://localhost:6080/oauth2/token",
			AuthStyle: oauth2.AuthStyleInHeader,
		},
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	app.Get("/internal-login", func(ctx *fiber.Ctx) error {
		url := conf.AuthCodeURL("")

		return ctx.Render("login", fiber.Map{
			"authorize_url": url,
		})
	})

	app.Get("/oauth2/callback", func(ctx *fiber.Ctx) error {
		code := ctx.Query("code")
		tok, _ := conf.Exchange(ctx.UserContext(), code)

		fmt.Println(tok)

		return nil
	})

	_ = app.Listen(os.Getenv("BLOG_ADDR"))
}
