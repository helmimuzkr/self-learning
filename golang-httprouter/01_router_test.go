package golang_httprouter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
)

// exmaple handler for Httprouter
func indexHttpRouter(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// code here
	fmt.Println("HttpRouter")
}

// example handler for http.Servemux
func indexMux(w http.ResponseWriter, r *http.Request) {
	// code here
	fmt.Println("ServeMuxl")
}

func TestRouter(t *testing.T) {
	// http.Servemux vs httprouter

	// Membuat objeknya
	mux := new(http.ServeMux)
	// Buat objek router dengan memanggil construct dari httprotuer
	router := httprouter.New()


	// Registrasi routernya
	mux.HandleFunc("/", indexMux)
	router.GET("/", indexHttpRouter) // sama seperti router.Handle(http.MethodGet, "/", indexHttpRouter)

	// Karena httprouter.R	outer implementasi dari http.Handler juga, maka dia bisa dipass sebagai handler pada http.ListenAndServe
	http.ListenAndServe("localhost:8000", mux)
	http.ListenAndServe("localhost:8000", router)
}