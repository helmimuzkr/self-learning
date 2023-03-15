package database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestQuerySqlComplex(t *testing.T) {
	// open
	db := GetConnection()
	defer db.Close()

	// context
	ctx := context.Background()

	// query
	query := "SELECT id, name, email, balance, created_at, birth_date, married FROM customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var (
			id int32
			name	string
			// GO tidak support data dengan value Null. Ketika ada column yang bisa memiliki value/defaultnya Null, maka akan error.
			// Maka dari itu tipe data email, balance, dan birth_date berbeda, karena menggunakan tipe data struct dari package database/sql.
			// Untuk pendeklarasiannya berbeda-beda, tergantung tipe data yang dari column databasenya.
			// tipe data struct ini didalamnya terdapat 2 field, yaitu field Tipe Data dan field Valid
			// contoh : 
			// Pada variabel email dengan tipe data struct sql.NullString, di dalamnya terdapat field String dengan tipe data string
			// dan field Valid dengan tipe data boolean
			email	sql.NullString
			balance sql.NullInt32
			birth_date	sql.NullTime
			// Akan terjadi error ketika belum dilakukan parsing ke time.Time. 
			// Karena saat driver mysql GO melakukan query ke database akan menghasilkan []byte atau []uint8. 
			// Agar bisa diterima dan dilakukan Scan harus di parse ke time.Time dulu agar menjadi string.
			// Caranya, menambahkan parameter parseDate=true di method Open connection, agar parsing otomatis ke time.Time.
			// yang tadinya sql.Open(nama_driver, "username:password@tcp(host:port)/database_name")
			// menjadi => sql.Open(nama_driver, "username:password@tcp(host:port)/database_name?parseTime=true")
			created_at	time.Time
			married bool
		)

		err := rows.Scan(&id, &name, &email, &balance, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("=====================================================")
		fmt.Println("Id \t\t: ", id)

		fmt.Println("Name \t\t: ", name)

		if email.Valid {
			fmt.Println("Email \t\t: ", email)
		}else {fmt.Println("Email \t\t: -")}

		if balance.Valid {
			fmt.Println("Balance \t: ", balance)
		}else {fmt.Println("Balance \t: -" )}
		
		fmt.Println("Created at \t: ", created_at)

		if birth_date.Valid{
			fmt.Println("Birth date \t: ", birth_date)
		}else {fmt.Println("Birth date \t: -")}

		fmt.Println("Married \t: ", married)
	}
}