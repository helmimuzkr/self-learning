package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func formPostHandler(w http.ResponseWriter, r *http.Request) {
	// Parsing data form dari request agar bisa diolah
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// Ambil data yang sudah diparsing lalu simpan ke dalam variabel name.
	// name ini nantinya akan menyimpan data url.Values(tipe Map) dengan key Name.
	name := r.PostForm["name"]
	fmt.Fprint(w, name)
}

func TestPostForm(t *testing.T) {
	// Membuat data dummy yang akan dikirimkan sebagai request'
	requestBody := strings.NewReader("name=Helmi&name=Muzakir")
	
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	// Buat content type aplication/x-www-form-urlencode
	// Content type ini wajib ada karena kita akan mencoba mengirim form post
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	
	recorder := httptest.NewRecorder()

	formPostHandler(recorder, request)

	// Melihat isi response body
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body)) // [Helmi Muzakir] = slice string
}