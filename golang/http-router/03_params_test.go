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

// Params
// httprouter.Handle memiliki parameter ketiga yaitu httprouter.Params.
// httprouter.Params merupakan tempat menyimpan parameter dari client.

// Berbeda dengan query parameter di dalam request, httprouter.Params ini bisa membuat sebuah parameter yang dinamis
// Misalnya :
// httprouter.Params: localhost:8080/genres/1/movies/2
// http.Reuqest: localhost:8080/?genres_id=1
// Pada Params setelah parameter bisa ditambahkan path lalu ditambahkan lagi parameter, sedangkant query parameter pada http.Request tidak bisa.

// Untuk menandai bahwa url tersebut adalah sebuah parameter yaitu
// dengan cara menambahkan simbol ":" pada awal kata saat pendaftaran endpoint
// Contoh "genres/:id"

// Parameter yang sudah ditandai saat register endpoint akan dikumpulkan lalu disimpan kedalam httprouter.Params.

func TestParams(t *testing.T) {
	// Buat objek router dengan memanggil construct dari httprotuer
	router := httprouter.New()

	// registrasi endpoint
	router.GET("/genres/:genreId/movies/:movieId", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		genreId := params.ByName("genreId")
		movieId := params.ByName("movieId")
		
		text := fmt.Sprintf("genre id: %s, movie id: %s", genreId, movieId)
		w.Write([]byte(text))
	})
	
	// buat request dan recorder untuk test
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/genres/1/movies/2", nil)
	recorder := httptest.NewRecorder()

	// serve http
	router.ServeHTTP(recorder, request)

	// Ambil response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Lakukan test
	assert.Equal(t, "genre id: 1, movie id: 2", string(body))
}