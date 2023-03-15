package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// queryParameter handler
func queryParameter(w http.ResponseWriter, r *http.Request) {
	// Query parameter digunakan untuk mengirim data dari client ke server. 
	// Query parameter ditempatkan pada URL
	// URL melakukan parsing query parameter dan menyimpannya dalam bentuk map[string][]string
	// Format query parameter => https://domain.com/endpoint?nama_query_parameter=value
	id := r.URL.Query().Get("id")

	fmt.Fprintf(w, "value pada query parameter ID : %s", id)
}
func TestQueryParameter(t *testing.T) {
	// make a request with get method
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/test?id=1", nil)
	// make a recorder for response
	recorder := httptest.NewRecorder()

	// call handler and passing recorder and request
	queryParameter(recorder, request)

	// check response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyStr := string(body) // byte to string

	fmt.Println(bodyStr)
}

// multipleParameter handler
func multipleParameter(w http.ResponseWriter, r *http.Request) {
	var query url.Values = r.URL.Query() // akan return map[string][]string dengan alias Values
	// ambil query dengan key "name" dalam values
	names := query["name"] // return slice string dari key parameter "name"
	// kenapa tidak menggunakan method get? karena method get hanya mengembalikan 1 data
	// yaitu index pertama dari return slice string dari key parameter "name".

	// kalau bingung kenapa Custom Type(Values) bisa dijadikan reciever value pada method
	// baca ini => Custom Type and Receiver Functions. https://articles.wesionary.team/slices-custom-types-and-receiver-functions-methods-in-golang-cdce4c01a5e8

	// join slice names menjadi string
	name := strings.Join(names, " ")
	fmt.Fprintf(w, "value pada query parameter name : %s", name)
}
func TestMultipleParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/data?name=Helmi&name=Muzakir", nil)
	recorder := httptest.NewRecorder()

	multipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyStr := string(body) // byte to string
	fmt.Println(bodyStr)
}