package golang_httprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// ● Apa yang terjadi jika terjadi panic pada logic Handler yang kita buat?
// Secara otomatis akan terjadi error, dan web akan berhenti mengembalikan response
// ● Kadang saat terjadi panic, kita ingin melakukan sesuatu, misal memberitahu jika terjadi kesalahan
// di web, atau bahkan mengirim informasi log kesalahan yang terjadi
// ● Sebelumnya, seperti yang sudah kita bahas di materi Go-Lang Web, jika kita ingin menangani panic,
// kita harus membuat Middleware khusus secara manual
// ● Namun di Router, sudah disediakan untuk menangani panic, caranya dengan menggunakan
// attribute PanicHandler : func(http.ResponseWriter, *http.Request, interface{})

func TestPanic(t *testing.T) {
	// Buat objek router dengan memanggil construct dari httprotuer
	router := httprouter.New()

	// Buat panic handlernya
	// Assign data anonymous func dengan parameter http.ResponseWriter, http.Reuqest, dan interface
	// ke dalam field/property panic handler
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "error: ", i)
	}

	// Registrasi routernya
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("502")
	})

	// Buaat objek untuk keperluan test
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	router.ServeHTTP(recorder, request)

	// Ambil body responsenya
	body, _ := io.ReadAll(recorder.Result().Body)

	// Lakukan test
	assert.Equal(t, "error: 502", string(body))
}