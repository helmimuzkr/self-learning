package golang_httprouter

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	// Buat objek router dengan memanggil construct dari httprotuer
	router := httprouter.New()

	// registrasi endpoint dengan method get
	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Write([]byte("Test Handle"))
	})

	// buat request dan recorder testnya
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/", nil)
	recorder := httptest.NewRecorder()

	// jalankan http router
	router.ServeHTTP(recorder, request)

	// ambil response dari hasil record
	// Lalu gunakan io.ReadAll untuk membaca response
	// Hasil akhir akan menjadi slice of bytes
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Lakukan uji coba dengan
	// Arg pertama adalah Assert Testign
	// Arg kedua adalah expectation
	// Arg ketiga adalah actual result dari code diatas.
	assert.Equal(t, "Test Handle", string(body))
}