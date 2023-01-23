package database

import (
	"context"
	"fmt"
	"testing"
)

// QueryContext
func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer" 
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close() // tutup rows kalau func udah selesai dieksekusi

	var data DataCustomer
	for rows.Next() {
		// deklarasi variabel untuk penampungan data dari database
		var id int32
		var name string
		
		// Pada method Scan ini di dalamnya akan ada proses assignment value 
		// ke variabel pointer(variabel yang sudah dideklarasikan), jadi argsnya harus pointer.
		// Hanya mengembalikan error.
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		// simpan data id, name ke dalam struct Customer
		customer := &Customer{
			Id: id,
			Name: name,
		}

		// tambah data customer ke struct DataCustomer
		data.getData(customer)
	}

	fmt.Println(data)
} 