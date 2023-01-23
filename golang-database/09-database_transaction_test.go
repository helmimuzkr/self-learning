package database

import (
	"context"
	"database/sql"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabaseTransaction(t *testing.T) {
	// open database connection
	dataSource := "root:@tcp(localhost:3306)/belajar_golang?parseTime=true"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// conf database connection
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	// context
	ctx := context.Background()

	// Get a tx for making transaction request
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// prepare statement (insert data)
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO comments(name, comment) VALUES(?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()


	// for-loop to creating and insert dummy data
	for i := 16; i <= 20; i++ {
		// creating dummy input data for subtitution to parameter ExecContext
		name := "helmi"
		comment := "Komentar ke-" + strconv.Itoa(i)

		// exec query
		_, err := stmt.ExecContext(ctx, name, comment)
		if err != nil {
			panic(err)
		}
	}

	// Data masih belum melakukan update sampai dilakukan tx.Commit atau tx.Rollback
	// bisa dilihat tablenya dari terminal lain

	// data akan di update
	// tx.Commit()

	// data akan di kembalikan seperti sebelum Transaction dimulai
	// tx.Rollback()
}