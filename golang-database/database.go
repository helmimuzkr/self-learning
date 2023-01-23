package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB{
	// open connection or database pooling(management connection)
	// nama_driver, "username:password@tcp(host:port)/database_name"
	// parseTime = true digunakan untuk parsing saat query ke database
	// dari tipe data DATE, DATETIME, TIMESTAMP ke time.Time(string)
	dataSource := "root:@tcp(localhost:3306)/belajar_golang?parseTime=true"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error()) // kalau connection gagal, langsung panic error
	}
	// defer db.Close() // menutup database kalau sudah digunakan

	// fmt.Println("Connection Successfuly")

	// set koneksi yang nganggur sebanyak 10
	db.SetMaxIdleConns(10)
	// set maksimal koneksi yang bisa dipake
	db.SetMaxOpenConns(100)
	// set waktu nganggur koneksi, kalau udah ngelewatin bakalan ditutup
	db.SetConnMaxIdleTime(5 * time.Minute)
	// set waktu maksimal koneksi
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
