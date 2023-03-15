package golang_httprouter

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// ● Selain panic handler, Router juga memiliki not found handler
// ● Not found handler adalah handler yang dieksekusi ketika client mencoba melakukan request URL
// yang memang tidak terdapat di Router
// ● Secara default, jika tidak ada route tidak ditemukan, Router akan melanjutkan request ke
// http.NotFound, namun kita bisa mengubah nya
// ● Caranya dengan mengubah router.NotFound = http.Handler

func TestNotFound(t *testing.T) {
	// Buat objek router dengan memanggil construct dari httprotuer
	router := httprouter.New()

	// Assign value dengan tipe datapa http.Handler
	// yang nantinya akan dijalankan ketika url path tidak terdaftar dalam router endpoint
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("404: Page not found"))
	})

	// Karena kita ingin mencoba panic handler, tidak usah meregistrasi endpoint supaya terjadi error.
	// Apapun url pathnya akan terjadi error, karena path routenya belum diregistrasi

	// Buaat objek untuk keperluan test
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	router.ServeHTTP(recorder, request)

	// Ambil body responsenya
	body, _ := io.ReadAll(recorder.Result().Body)

	// Lakukan test
	assert.Equal(t, "404: Page not found", string(body))
}