package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	// Membuat handler dengan type data http.HandlerFunc.
	// Type data ini menerima anonymous func dengan parameter htt.ResponseWriter dan http.Requerst.
	// Kenapa menggunakan HandlerFunc? karena pada saat memasukan value ke field Handler di struct
	// Server hanya menerima value dengan type data interface Handle.
	// Type data HandlerFunc termasuk implementasi dari interface Handle.
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// Logic Web
		fmt.Fprint(writer, "Hello World")
	}
	
	// membuat server
	server := http.Server{
		Addr: "localhost:8080", // localhost address
		Handler: handler, // handle interface, untuk input request dan menerima response
	}
	// eksekusi server, nanti akan return error
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}