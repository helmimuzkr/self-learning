package database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConection (t *testing.T) {
	// open connection or database pooling(management connection)
	// sql.Open(nama_driver, "username:password@tcp(host:port)/database_name")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang")
	if err != nil {
		panic(err.Error()) // kalau connection gagal, langsung panic error
	}
	defer db.Close() // menutup database kalau sudah digunakan


}
