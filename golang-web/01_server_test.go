package web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// membuat server
	server := http.Server{
		Addr: "localhost:8080", // localhost address
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}