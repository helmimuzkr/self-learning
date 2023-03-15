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

// ● Saat menggunakan ServeMux, kita tidak bisa menentukan HTTP Method apa yang digunakan
// untuk Handler
// ● Namun pada HttpRouter, kita bisa menentukan HTTP Method yang ingin kita gunakan, lantas apa yang
// terjadi jika client tidak mengirim HTTP Method sesuai dengan yang kita tentukan?
// ● Maka akan terjadi error Method Not Allowed
// ● Secara default, jika terjadi error seperti ini, maka Router akan memanggil function http.Error
// ● Jika kita ingin mengubahnya, kita bisa gunakan httprouter.MethodNotAllowed = http.Handler

func TestMethodNotAllowed(t *testing.T) {
	// Buat objek router dengan memanggil construct dari httprotuer
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Method not allowed"))
	})

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("Hello World"))
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "http://localhost:8080/", nil)
	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
	
	assert.Equal(t, "Method not allowed", string(body))
}