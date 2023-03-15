package config

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	//  open connection
	db, err := sql.Open(
		"mysql", // driver
		"root:@tcp(localhost:3306)/belajar_golang?parseTime=true", // data source
	)
	if err != nil {
		panic(err)
	}

	// conf database pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

