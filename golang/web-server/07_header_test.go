package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Request Header
func requestHeader(w http.ResponseWriter, r *http.Request) {
	// Header juga disimpan dengan tipe map[string][]string
	// Menyimpan value yang diget dari key "content-type"
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}
func TestHeaderRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// menambah header baru dengan parameter (key, value)
	request.Header.Add("content-type", "aplication/json") 
	
	recorder := httptest.NewRecorder()

	requestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	
	fmt.Println(string(body)) // "aplication/json"
}

// Add new Response Header
func responseHeader(w http.ResponseWriter, r *http.Request) {
	// tambah header
	w.Header().Add("MZKR-Powered-By", "Helmi")

	fmt.Fprint(w, "OK")
}
func TestHeaderResponse(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	responseHeader(recorder, request)

	// mengambil header response yang sudah dibuat di handle dengan menggunakan
	// Method get Header yang tersedia di struct *httptest.ResponseRecorder, mengembalikan struct http.Header
	// dan, Method Get dengan parameter key yang nantinya akan mengembalikan value(string)
	// Pada saat pemganggilan header response tidak Case Sensitive
	response := recorder.Header().Get("mzkr-powered-by")
	fmt.Println(response)
}