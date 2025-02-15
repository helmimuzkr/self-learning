package main

import (
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("incoming request", "host", r.URL.Host, "url", r.URL.RawPath)
		w.Write([]byte("Hello from app1"))
	})

	port := ":8081"
	slog.Info("Listening on", "port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}
