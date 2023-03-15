package web

import (
	"fmt"
	"net/http"
	"testing"
)

// Membuat handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Method) // request method, lalu diterima writer
	fmt.Fprint(w, r.RequestURI) // requst url, lalu diterima writer
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Method) // request method, lalu diterima writer
	fmt.Fprint(w, r.RequestURI) // requst url, lalu diterima writer
}

func TestRequest(t *testing.T) {
	// membuat mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/about", AboutHandler)

	// membuat server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}