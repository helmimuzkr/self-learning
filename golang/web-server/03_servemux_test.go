package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	// HandlerFunc tidak bisa digunakan untuk membuat banyak endpoint URL, karena itu dibutuhkan ServeMux.
	// ServeMux adalah implementasi dari Handler. 
	// ServerMux digunakan untuk menggabungkan kumpulan handle-handle lainnya
	mux := http.NewServeMux() 	// Membuat serve mux
	
	// disini menggunakan method HandleFunc dari mux.
	// mux.HandleFunc(endpoint, handler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Home")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "About")
	})

	// Membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
if err != nil {
	panic(err)
}
}