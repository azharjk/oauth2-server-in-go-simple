package main

import (
	"github.com/estradax/exater/cmd/account/handler"
	"github.com/gofiber/fiber/v2"
)

func setupAuthRoute(app *fiber.App) {
	app.Get("/register", handler.Unauthorized, handler.Register).Name("register")
	app.Post("/register", handler.Unauthorized, handler.RegisterUser).Name("registerUser")

	app.Get("/login", handler.Unauthorized, handler.Login).Name("login")
	app.Post("/login", handler.Unauthorized, handler.CreateSession).Name("createSession")

	app.Post("/logout", handler.Authorized, handler.Logout).Name("logout")
}
