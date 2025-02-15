package main

import (
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("incoming request", "addr", r.RemoteAddr)
		w.Write([]byte("App 2"))
	})

	port := ":8082"
	slog.Info("Listening on", "port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}
