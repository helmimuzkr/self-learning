package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func main() {
	e := echo.New()

	e.GET("/login", loginHandler)
	e.GET("/callback", callbackHandler)
	e.GET("/calendar", calendarHandler)

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}

func loginHandler(c echo.Context) error {
	conf := &oauth2.Config{
		RedirectURL:  "http://localhost:8000/callback",
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}
	state := "random_state"

	url := conf.AuthCodeURL(state)

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["state"] = state
	sess.Values["conf"] = conf
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func callbackHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)

	state := sess.Values["state"]
	conf := sess.Values["conf"].(*oauth2.Config)

	if state != c.FormValue("state") {
		log.Println("state is not valid")
		return c.JSON(400, "state is not valid")
	}

	ctx := context.Background()

	token, err := conf.Exchange(ctx, c.FormValue("code"))
	if err != nil {
		return c.JSON(400, err.Error())
	}



	client := conf.Client(ctx, token)
	sess.Values["client"] = client
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8000/calendar")
}

func calendarHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)

	client := sess.Values["client"].(*http.Client)

	calendarService, err := calendar.New(client)
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
