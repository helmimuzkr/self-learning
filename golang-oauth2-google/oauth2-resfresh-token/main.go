package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func main() {
	NewConfOauth2()

	token, err := GetTokenFromFile("token.json")
	if err != nil {
		log.Println(err)
		// file not exist, then get token from web
		newToken, err := GetTokenFromWeb()
		if err != nil {
			return
		}
		token = newToken

		// Saves a token to a file "token.json"
		SaveToken("token.json", token)
	}

	// refresh token using old token
	tokenSource := Conf.TokenSource(context.TODO(), token)
	// get a new token
	newToken, err := tokenSource.Token()
	if err != nil {
		log.Fatal(err)
	}

	// update token
	// check if new token not same as the old token
	if newToken.AccessToken != token.AccessToken {
		// save new token to a file "token.json"
		SaveToken("token.json", newToken)
		log.Println("Saved new token:", newToken.AccessToken)
	}

	fmt.Println("NEW TOKEN === ")
	fmt.Println("Access: ", newToken.AccessToken)   // ya29.a0AVvZVsr4FnqugD70It6vxE3nkJPKOOXGVvTlya-i3KJQBwjZDKSSMow-1BLu9LBRpuRgf-J_hDqFKlHUL6lQ5INyiD0Zm6Bk27OXZA8TO_l-d5QFgkVg4UOEhOvMCUlitwWUkmSHnIz2fg7Uy1FCWrYAHZBL9BsaCgYKAacSAQASFQGbdwaIKadsoSpgHbUgpdbBk1ss_w0166
	fmt.Println("Refresh: ", newToken.RefreshToken) // 1//0gY2pudM1N7HwCgYIARAAGBASNwF-L9Ir890RUH-kvDwDVvuZh8C9viWvtkCDtkZBINE0B5vdcZNSAvdXereGHGVn03AiShUVv9U
	fmt.Println("Expiry: ", newToken.Expiry)        // 2023-02-17 23:25:40.304021893 +0700 WIB
	fmt.Println()
	fmt.Println("OLD TOKEN === ")
	fmt.Println("Access: ", token.AccessToken)   // ya29.a0AVvZVsr4FnqugD70It6vxE3nkJPKOOXGVvTlya-i3KJQBwjZDKSSMow-1BLu9LBRpuRgf-J_hDqFKlHUL6lQ5INyiD0Zm6Bk27OXZA8TO_l-d5QFgkVg4UOEhOvMCUlitwWUkmSHnIz2fg7Uy1FCWrYAHZBL9BsaCgYKAacSAQASFQGbdwaIKadsoSpgHbUgpdbBk1ss_w0166
	fmt.Println("Refresh: ", token.RefreshToken) // 1//0gY2pudM1N7HwCgYIARAAGBASNwF-L9Ir890RUH-kvDwDVvuZh8C9viWvtkCDtkZBINE0B5vdcZNSAvdXereGHGVn03AiShUVv9U
	fmt.Println("Expiry: ", token.Expiry)        // 2023-02-17 23:25:40.304021893 +0700 WIB

	fmt.Println()
	client := oauth2.NewClient(context.TODO(), tokenSource)

	srv, err := calendar.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}

}
