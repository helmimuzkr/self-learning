package utill

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func InitPostgresDB() {
	if DB != nil {
		return
	}

	c := initConfig()

	source := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.HOST, c.PORT, c.USER, c.PASSWORD, c.DBNAME)

	db, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatal("database connection error : ", err.Error())
		return
	}

	DB = db
}

// Config env
type PostgresConfig struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DBNAME   string
}

func initConfig() *PostgresConfig {
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
		return nil
	}

	c := new(PostgresConfig)
	if err := viper.Unmarshal(c); err != nil {
		log.Fatal(err)
		return nil
	}

	slog.Info("config info", "config", *c)

	return c
}
