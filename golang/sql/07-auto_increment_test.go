package database

import (
	"context"
	"fmt"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments(name, comment) VALUES(?, ?)"
	
	// Result adalah object yang dikembalikan ketika kita menggunakan function Exec
	result, err := db.ExecContext(ctx, query, "helmi", "test")
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert data successfully")

	// Function (Result) LastInsertId() untuk mendapatkan Id terakhir yang dibuat secara auto increment 
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Last comment by id", lastInsertId)
}