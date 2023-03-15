package main

import (
	"log"

	"github.com/spf13/viper"
)

// Config env
type Config struct {
	ClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	ClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	RedirectURL  string `mapstructure:"GOOGLE_REDIRECT_CALLBACK"`
}

func initConfig() *Config {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
		return nil
	}

	c := new(Config)
	if err := viper.Unmarshal(c); err != nil {
		log.Fatal(err)
		return nil
	}

	return c
}
