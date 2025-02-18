package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", loggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("name") == "" {
			http.Error(w, "missing in parameter query", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello %s!", r.URL.Query().Get("name"))
	}))

	port := ":8081"
	slog.Info("listening..", "port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("incoming request", "host", r.Host, "method", r.Method, "url", r.URL)

		respWrapper := NewResponseWrapper(w)
		next(respWrapper, r)

		slog.Info("prosses done", "status", respWrapper.statusCode, "message", string(respWrapper.body))
	}
}

// ResponseWriter wrapper to capture status and body
type responseWrapper struct {
	w http.ResponseWriter
	// store the status code manually instead of retrieving it with w.Header().Get() is
	// that http.ResponseWriter does not store the status code in the headers.
	// once w.WriteHeader() is called, the status is written to the connection, but Go does not provide a way to retrieve it.
	statusCode int
	// for simplicty the body intended for a message with a string type.
	body []byte
}

func NewResponseWrapper(w http.ResponseWriter) *responseWrapper {
	return &responseWrapper{w: w, body: []byte{}}
}

func (rw *responseWrapper) Write(message []byte) (int, error) {
	rw.body = message
	return rw.w.Write(message)
}

func (rw *responseWrapper) Header() http.Header {
	return rw.w.Header()
}

func (rw *responseWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.w.WriteHeader(statusCode)
}
