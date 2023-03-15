package warming_up

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getConnection() *sql.DB{
		// Open db
		db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang?parseTime=true")
		if err != nil {
			panic(err)	
		}
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
		db.SetConnMaxIdleTime(5 * time.Minute)
		db.SetConnMaxLifetime(1 * time.Hour)

		return db
}

// Struct untuk menampung data
type Customer struct {
	Name string
	Balance sql.NullInt32
	Created_at sql.NullTime
}

func getCustomer() *[]Customer {
		// Get the connection then will returning a struct pointer, *sql.DB
		db := getConnection()

		// Assign context.Context to ctx
		ctx := context.Background()
	
		// Assigning a sql.Rows and error to rows and err
		rows, err := db.QueryContext(ctx, "SELECT name, balance, created_at FROM customer")
		if err != nil {
			panic(err)
		}
		defer rows.Close()	
		

		customers := &[]Customer{}
	
		// Mengambil data dengan melakukan perulangan
		for rows.Next() {
			customer := Customer{}
			rows.Scan(&customer.Name, &customer.Balance, &customer.Created_at)
			*customers = append(*customers, customer)
		}


		fmt.Println(&customers)
		fmt.Println("")
	
		return customers
}

func TestTypeColumn(t *testing.T) {
	customers := getCustomer()

	fmt.Println(&customers)
/* 
	for _, v := range *customers {
		fmt.Println(&v)
	} */
}
