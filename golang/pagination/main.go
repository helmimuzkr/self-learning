package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Worker struct {
	ID         int
	Name       string
	Department string
}

type Pagination struct {
	Page         int
	Limit        int
	Offset       int
	TotalPages   int
	TotalRecords int64
}

type Response struct {
	Data       interface{}
	Pagination Pagination
}

func main() {
	dsn := "root:helmi@tcp(localhost:3306)/pagination?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	// New echo
	e := echo.New()

	// Handler
	GetAllHandler := func(c echo.Context) error {
		// Get the query parameter
		strPage := c.QueryParam("page")
		page, _ := strconv.Atoi(strPage)

		// Get Pagination
		pagination := getPagination(db, page)

		// Get data worker
		workers := getWorkers(db, pagination)

		// Response
		response := Response{
			Data:       workers,
			Pagination: pagination,
		}

		return c.JSON(http.StatusOK, response)
	}

	e.GET("/", GetAllHandler)

	// server start
	if err := e.Start(":8000"); err != nil {
		fmt.Println(err)
		return
	}

}

func getWorkers(db *gorm.DB, pagination Pagination) []Worker {
	var listWorker []Worker
	db.Raw("SELECT * FROM Worker LIMIT ? OFFSET ?", pagination.Limit, pagination.Offset).Find(&listWorker)

	return listWorker
}

func getPagination(db *gorm.DB, page int) Pagination {

	// get total record
	var totalRecords int64
	db.Table("Worker").Count(&totalRecords)

	// Total Page
	totalPages := totalRecords / 10

	if page < 1 {
		page = 1
	}

	if page > int(totalPages) {
		page = int(totalPages)
	}

	// Limit
	limit := 10

	// Calculate offset
	offset := (page - 1) * limit

	pagination := Pagination{
		Page:         page,
		Limit:        limit,
		Offset:       offset,
		TotalPages:   int(totalPages),
		TotalRecords: totalRecords,
	}

	return pagination
}
