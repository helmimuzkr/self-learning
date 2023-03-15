package database

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type Customer struct {
	Id          int32
	Name, Email string
	Balance     int32
	Created_at  time.Time
	Birth_date 	time.Time
	Married		bool
}

type DataCustomer struct {
	Customers []Customer
}

func (dc *DataCustomer) getData(c *Customer) {
	fmt.Println("memory di dalam func", &c)
	dc.Customers = append(dc.Customers, *c)
}

// ExecContext
func TestExecSql(t *testing.T) {
	// get connection
	db := GetConnection()
	defer db.Close() // tutup ketika function testExec selesai dieksekusi

	// init context
	ctx := context.Background()

	// query sql
	query := "INSERT INTO customer(name) VALUES('helmi')"

	// execution sql query and returning an error
	// _(result), err := db.ExectConext(context, query, ...[]args)
	// ExecContext tidak mengembalikan result/data
	// Biasa digunakan untuk Create, Update, dan Delete
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data")
}
