package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type webResponse map[string]interface{}

// -------------------
// Redirect to consent screen
// -------------------
func authGoogle(ctx echo.Context) error {
	// If you want to refresh the token.
	// But, you need to add adtional parameter option in AuthCodeURL to get the token refresh.
	// GoogleConf.AuthCodeURL(State, oauth2.AccessTypeOffline).
	return ctx.Redirect(http.StatusTemporaryRedirect, GoogleConf.AuthCodeURL(State))
}

// -------------------
// Auth callback
// -------------------
func authGoogleCallback(ctx echo.Context) error {
	if ctx.QueryParam("state") != State {
		return ctx.JSON(401, webResponse{"message": "state not valid"})
	}

	code := ctx.QueryParam("code")

	token, err := GoogleConf.Exchange(ctx.Request().Context(), code)
	if err != nil {
		log.Println(err)
		return ctx.JSON(500, webResponse{"message": "internal server error"})
	}

	// Refresh the token using this syntax.
	// tokenSource := GoogleConf.TokenSource(ctx.Request().Context(), token)
	// newToken := tokenSource.Token()

	// Create request to google api for getting user info.
	baseURL := "https://www.googleapis.com/oauth2/v1/userinfo"
	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		log.Println(err)
		return ctx.JSON(500, webResponse{"message": "internal server error"})
	}
	bearer := "Bearer " + token.AccessToken
	request.Header.Set("Authorization", bearer)
	// Do sends an HTTP request and returns an HTTP response.
	client := http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return ctx.JSON(500, webResponse{"message": "internal server error"})
	}
	// Hook the response payload.
	// Decode from stream byte to userPayload using json Decoder
	userPayload := make(map[string]interface{})
	if err := json.NewDecoder(response.Body).Decode(&userPayload); err != nil {
		log.Println(err)
		return ctx.JSON(500, webResponse{"message": "internal server error"})
	}

	return ctx.JSON(200, webResponse{
		"message": "ok",
		"data":    userPayload,
	})
}
