package golang_httprouter

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed src
var srcDir embed.FS

func TestServeFile(t *testing.T) {
	// Buat objek router dengan memanggil construct dari httprotuer
	router := httprouter.New()

	// Masuk 1 folder kedalam srcDir
	srcDir, err := fs.Sub(srcDir, "src")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Kalau ga masuk 1 folder nanti di path urlnya harus disertakan dengan nama directory pennyimpanannya
	// Contoh
	// Tanpa sub directory: 
	// route path => "/file/src/*filepath"
	// url path => "http://localhost:8080/file/src/greetings.txt"

	// registrasi route endpoint  dan serve file
	router.ServeFiles("/file/*filepath", http.FS(srcDir))

	// buat keeprluan test
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/file/greetings.txt", nil)
	router.ServeHTTP(recorder, request)

	// ambil response body
	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

	// Lakukan test
	assert.Equal(t, "Hallo Helmi!", string(body))
}