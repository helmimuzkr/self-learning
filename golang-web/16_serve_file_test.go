package web

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
)

func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil link URLnya
	url := r.URL.RequestURI()
	// Split URL dengan separator "/", dan akan menjadi slice
	sliceURL := strings.Split(url, "/")
	// Ambil element terakhir/filename
	filename := sliceURL[len(sliceURL)-1]

	if filename != "" {
		// if not empty, then ...
		http.ServeFile(w, r, "./resource/img/" + filename)
	}else {
		fmt.Fprint(w, "Not found")
	}
}

func TestServeFile(t *testing.T) {
	// route
	http.HandleFunc("/", ServeFileHandler)

	// addr
	addr := "localhost:8080"
	log.Println("running on server ", addr)

	// running server
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}