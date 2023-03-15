package database

import (
	"context"
	"fmt"
	"testing"
)

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// uname := "admin"
	// pwd := "admin"
	
	// Jika ada user yang ingin memanfaatkan celah menggunakan SQL Injection,
	// maka hal yang dilakukan akan memanipulasi input username seperti dibawah ini.
	uname := "admin'; #"
	pwd := "salah"
	
	query := "SELECT name, username FROM users WHERE username='" + uname + 
						"' AND password='" + pwd + "' LIMIT 1"
	
	// Kalau diprint maka querynya akan seperti ini
	fmt.Println(query)
	// SELECT name, username FROM users WHERE username='admin'; #' AND password='salah' LIMIT 1
	// artinya, script querynya cuman sampai where username='admin'.
	// lalu ditutup dengan simbol => ; 
	// dan script setelahnya dijadikan komentar dengan simbol komentar pada query sql yaitu => #
	// karena script setelah admin diabaikan, maka akan menselect berdasarkan ID saja, tanpa password.
	
	// Agar saat query ke database dan simbol-simbol seperti tadi diabaikan, golang menyediakan fitur 
	// sql dengan parameter/argumen.

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
		
	if rows.Next() {
		var (
			name, uname string
		)

		err := rows.Scan(&name, &uname)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Nama\t\t: %s\nUsername\t: %s", name, uname)
	}
}