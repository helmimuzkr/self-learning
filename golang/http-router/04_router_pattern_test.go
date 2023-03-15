package golang_httprouter

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Named Parameter
// Sebelumnya kita sudah menggunakan Named parameter

// Pattern: /user/:user

//  /user/gordon              match
//  /user/you                 match
//  /user/gordon/profile      no match
//  /user/                    no match

// Berikut ini adalah contoh named parameter dengan 1 parameter.
func TestNamedParameter(t *testing.T) {
	// Buat objek router dengan memanggil construct dari httprotuer
	router := httprouter.New()
	
	// Buat handle
	namedHandler := func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		movieName := params.ByName("name")
		w.Write([]byte("Movie name: " + movieName))
	}

	// register endpoint dan membuat handle
	router.GET("/movie/:name", namedHandler)

	// Membuat request dan recorder untuk test
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/movie/jujutsu-kaisen", nil)
	recorder := httptest.NewRecorder()

	// Serving http dengan argumen dari recorder dan request yang sudah dibuat tadi
	router.ServeHTTP(recorder, request)

	// Ambil body response dari hasil recorder
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Lakukan test
	assert.Equal(t, "Movie name: jujutsu-kaisen", string(body))
}


// Catch  All Parameter
// Catch all parameter digunakan untuk menangkap semua parameter url setelah prefix atau pendefinisian catch parameter.

// Pattern: /src/*filepath

//  /src/                     match
//  /src/somefile.go          match
//  /src/subdir/somefile.go   match

// Untuk mendefinisikan kalau path url sebagai catch parameter, menggunakan simbol "*" pada awal kata
// Contoh : /file/*filename
// maka file/... kebelakang akan dianggap sebagai parameter.

// Misalnya /file/src/img/profile-picture.png
// yang akan diambil adalah /src/img/profile-picture.png
func TestCatchAllParameter(t *testing.T) {
	// Buat objek router dengan memanggil construct dari httprotuer
	router := httprouter.New()

	// Buat Handle
	catchHandler := func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		filename := params.ByName("filename")
		w.Write([]byte("Nama file: " + filename))
	}

	// Register endpoint
	router.GET("/file/*filename", catchHandler)

	// Buat request dan recorder untuk kebutuhan test
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8008/file/src/img/profile-picture.png", nil)

	// Serving http
	router.ServeHTTP(recorder, request)

	// Ambil response body dari  hasil recorder
	responseBody, _ := io.ReadAll(recorder.Result().Body)

	// Lakukan test
	assert.Equal(t, "Nama file: /src/img/profile-picture.png", string(responseBody))
}