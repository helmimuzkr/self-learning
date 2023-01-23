package database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments(name, comment) VALUES(?, ?)"
	// creating prepare statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	
	// using for-loop to make 10 data
	for i := 0; i < 10; i++ {
		name := "user" + strconv.Itoa(i)
		comment := "Komen dari user" + strconv.Itoa(i)

		// exec query with stmt
		result, errStmt := stmt.ExecContext(ctx, name, comment)
		if errStmt != nil {
			panic(errStmt)
		}

		// taking last insert id
		lastInserId, errLastInsertId := result.LastInsertId()
		if errLastInsertId != nil {
			panic(errLastInsertId)
		}

		fmt.Println("Comment from id : ", lastInserId)
	}
}