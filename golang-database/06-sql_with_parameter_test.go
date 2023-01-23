package database

import (
	"context"
	"fmt"
	"testing"
)

func TestSqlWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	uname := "admin'; #" // expect gagal
	pwd := "admin"
	
	// Untuk menandai sebuah SQL membutuhkan parameter, kita bisa gunakan karakter ? (tanda tanya)
	query := "SELECT name, username FROM users WHERE username = ? AND password = ? LIMIT 1"

	// lalu pada argumen ketiga, kita mengirim input user ke sql query
	// db.QueryContext(context, query, args ...interface{})

	// nantinya argumen ketiga akan disubstitusi ke query sql sesuai urutannya.
	// misal, di script query ada parameter yang membutuhkan input username = ? AND password = ?
	// maka pada argumen ke 3 diisi dengan variabel yang berisi input user untuk dikirimkan ke query sql yang ditandai dengan ?(tanda tanya)
	// contoh
	// db.QueryContext(arg1, arg2, ...args3) => db.QueryContext(arg1, arg2, input1, input2)
	// untuk input parameter sesuai urutan!
	rows, err := db.QueryContext(ctx, query, uname, pwd)
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

		fmt.Printf("Nama\t\t: %s\nUsername\t: %s\n", name, uname)
	}else {
		fmt.Println("Login failed") // expect gagal
	}
}