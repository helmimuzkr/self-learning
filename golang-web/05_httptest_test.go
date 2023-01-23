package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// handler
func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Index")
}

func TestHttpTest(t *testing.T) {
	// httptest.NewRequest adalah function yang digunakan untuk membuat http.Request, 
	// yang akan dikirim sebagai simulasi unit test
	// Untuk membuat testnya dengan format : httptest.Request(method, server, requestBody(bisa nil)) 
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	// httptest.NewRecorder adalah function yang digunakan untuk membuat http.ResponseWriter. 
	// Bentuknya struct bantuan untuk merekam HTTP response dari hasil testing yang telah dilakukan.
	// Untuk membuat testnya hanya dengan memanggil method  httptest.NewRecord()
	record := httptest.NewRecorder()

	// Panggil hadler dengan mengirim argumen record sebagai response, dan request sebagai request
	IndexHandler(record, request)

	// check respsonsenya
	response := record.Result() // return *http.Response
	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("status: ",response.Status)
	// respnose.Body membutuhkan io.ReadAll untuk membaca responsenya, lalu return []byte data dan err
	body, err := io.ReadAll(response.Body) 
	if err != nil {
		panic(err)
	}  
	// convert body to string
	bodyStr := string(body)
	fmt.Println("isi response: ", bodyStr)
}