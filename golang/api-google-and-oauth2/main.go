package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Oauth2 config instance
var GoogleConf *oauth2.Config

var State string = "random string"

func main() {
	e := echo.New()
	c := initConfig()

	GoogleConf = &oauth2.Config{
		RedirectURL:  c.RedirectURL,
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	e.GET("/auth/google", authGoogle)
	e.GET("/auth/google/callback", authGoogleCallback)

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
