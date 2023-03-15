package main

import (
	"log"
	"net/smtp"

	"github.com/spf13/viper"
)

func main() {
	log.Println("Setup env...")
	viper.SetConfigFile("./.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	config := make(map[string]interface{})
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	log.Println("Send mail...")
	username := config["smtp_username"].(string)
	password := config["smtp_password"].(string)
	host := config["smtp_host"].(string)
	port := config["smtp_port"].(string)
	smtpAddress := host + ":" + port
	auth := smtp.PlainAuth("", username, password, host)
	reciever := "campyukuser@gmail.com"
	body := "From: " + username + "\n" +
		"To: " + reciever + "\n" +
		"Subject: " + "verification social media email" + "\n\n" +
		"boyd messages"

	if err := smtp.SendMail(smtpAddress, auth, username, []string{reciever}, []byte(body)); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully sent mail")
}
