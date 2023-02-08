package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var (
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/callback",
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}
	RandomState = "random"
)

func main() {
	e := echo.New()

	e.GET("/login", loginHandler)
	e.GET("/callback", callbackHandler)

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}

func loginHandler(c echo.Context) error {
	url := GoogleOauthConfig.AuthCodeURL("random")
	// return http.Redirect(c.Response(), c.Request(), url, http.StatusTemporaryRedirect)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func callbackHandler(c echo.Context) error {
	if c.FormValue("state") != RandomState {
		log.Println("state is not valid")
		return c.JSON(400, "state is not valid")
	}

	ctx := context.Background()

	token, err := GoogleOauthConfig.Exchange(ctx, c.FormValue("code"))
	if err != nil {
		return c.JSON(400, err.Error())
	}

	// Get user
	// resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	// if err != nil {
	// 	return c.JSON(500, err.Error())
	// }
	// defer resp.Body.Close()

	// content, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return c.JSON(500, err.Error())
	// }

	// contentStr := string(content)

	calendarService, err := calendar.NewService(ctx, option.WithTokenSource(GoogleOauthConfig.Client(ctx, token)))
	if err != nil {
		return c.JSON(500, err.Error())
	}

	event := &calendar.Event{
		Summary:  "Camping",
		Location: "Bandung",
		Start: &calendar.EventDateTime{
			DateTime: time.Now().Format(time.RFC3339), // Wajib format RFC3339
			TimeZone: "Asia/Jakarta",
		},
		End: &calendar.EventDateTime{
			DateTime: time.Date(2023, 02, 10, 13, 20, 10, 10, time.Local).Format(time.RFC3339), // artinya YYYY-MM-DD HH-MM-SS-NS Location
			TimeZone: "Asia/Jakarta",
		},
		Attendees: []*calendar.EventAttendee{
			{Email: "sycgundul03@gmail.com"},
		},
	}

	event, err = calendarService.Events.Insert("primary", event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}

	return c.Redirect(http.StatusTemporaryRedirect, event.HtmlLink)
}
