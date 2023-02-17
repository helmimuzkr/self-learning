package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GOOGLE_CLIENT_ID     = "917459160831-udciilgtk29q50bs2t7j6fo71nebl3c7.apps.googleusercontent.com"
	GOOGLE_CLIENT_SECRET = "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"
	GOOGLE_REDIRECT_URL  = "http://localhost:8000/oauth/callback"

	Conf *oauth2.Config
)

func NewConfOauth2() {
	Conf = &oauth2.Config{
		ClientID:     GOOGLE_CLIENT_ID,
		ClientSecret: GOOGLE_CLIENT_SECRET,
		RedirectURL:  GOOGLE_REDIRECT_URL,
		Scopes: []string{
			"https://www.googleapis.com/auth/calendar",
			// "https://www.googleapis.com/auth/userinfo.profile",
			// "https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

// Request a token from the web, then returns the retrieved token.
func GetTokenFromWeb() (*oauth2.Token, error) {
	url := Conf.AuthCodeURL("random", oauth2.AccessTypeOffline)
	fmt.Printf("Click this link... \n%s\n", url)

	fmt.Println("Input in here...")
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Printf("Unable to read authorization code: %v", err)
		return nil, err
	}

	newToken, err := Conf.Exchange(context.Background(), authCode)
	if err != nil {
		log.Printf("Unable to retrieve token from web: %v", err)
		return nil, err
	}

	return newToken, nil
}

// Saves a token to a file path.
func SaveToken(filename string, token *oauth2.Token) error {
	log.Println("saving token to ", filename)

	// If the file doesn't exist, create it, or append to the file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// convert struct to json
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}

	// write data to file
	_, err = f.Write(data)
	if err != nil {
		return err
	}

	log.Println("successfully saving token")
	return nil
}

// retrieves a token from a local file.
func GetTokenFromFile(filename string) (*oauth2.Token, error) {
	log.Println("opening file from ", filename)

	// Open file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Stream decode using json
	token := &oauth2.Token{}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&token); err != nil {
		return nil, err
	}

	return token, nil
}
