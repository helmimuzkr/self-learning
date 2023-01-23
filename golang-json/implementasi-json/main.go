package main

import (
	"fmt"
	"net/http"

	"github.com/helmimuzkr/handler"
)

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/", handler.PostFormToJSON)

	addr := ":8080"
	fmt.Println("listening on port: ", addr)
	http.ListenAndServe("localhost"+addr, mux)
}
