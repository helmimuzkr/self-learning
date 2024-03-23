package handler

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/helmimuzkr/templ/types"
)

// In Go, once you've read the request body using io.ReadAll, the underlying r.Body stream is consumed,
// and when the second attempts to read from it(which is 'next' methods) will result in an empty stream or EOF error.
// So, the solution is, making duplicate reader for body request before executing the 'next' method.

func LogMiddleware(next types.HelHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()

		// buffer the request body
		var payload []byte
		if b, err := io.ReadAll(r.Body); err == nil {
			payload = b
		}
		r.Body.Close()

		// create a reader from the buffered payload
		r.Body = io.NopCloser(bytes.NewBuffer(payload))

		// create context.
		ctx := types.NewContext(w, r)

		// call the handler and log the error.
		err := next(ctx)

		slog.Log(context.Background(), slog.LevelInfo, "", "method", r.Method, "path", r.URL.Path, "took", time.Since(t).String(), "payload", string(payload), "error", err)
	}
}
