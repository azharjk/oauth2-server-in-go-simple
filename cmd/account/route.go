package main

import (
	"github.com/estradax/exater/cmd/account/handler"
	"github.com/estradax/exater/cmd/account/handler/oauth2"
	"github.com/gofiber/fiber/v2"
)

func setupAccountRoute(app *fiber.App) {
	app.Get("/", handler.Authorized, handler.Account).Name("account")
}

func setupAuthRoute(app *fiber.App) {
	app.Get("/register", handler.Unauthorized, handler.Register).Name("register")
	app.Post("/register", handler.Unauthorized, handler.RegisterUser).Name("registerUser")

	app.Get("/login", handler.Unauthorized, handler.Login).Name("login")
	app.Post("/login", handler.Unauthorized, handler.CreateSession).Name("createSession")

	app.Post("/logout", handler.Authorized, handler.Logout).Name("logout")
}

func setupOAuth2Route(app *fiber.App) {
	app.Get("/oauth2/authorize", handler.Authorized, oauth2.Authorize).Name("oauth2Authorize")
}
