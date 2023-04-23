package main

import (
	"flag"
	"fmt"
	"github.com/estradax/exater/internal/model"
	"github.com/estradax/exater/internal/model/client"
)

func main() {
	_ = model.Connect()

	redirectURI := flag.String("redirect_uri", "http://localhost:8080/oauth2/callback", "Redirect URI for new client.")
	flag.Parse()

	c := model.Client{
		Secret:      client.NewSecret(),
		Type:        "public",
		RedirectURI: *redirectURI,
	}

	_ = model.DB.Create(&c)

	fmt.Println("client id:", c.ID)
	fmt.Println("client secret:", c.Secret)
}
