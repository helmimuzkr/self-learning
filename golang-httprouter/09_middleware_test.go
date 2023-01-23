package golang_httprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Middleware with function
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware Auth : Ini perintah sebelum diteruskan ke handler selanjutnya")
		next.ServeHTTP(w, r)
	})
}
func ErrMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware err : Ini perintah sebelum diteruskan ke handler selanjutnya")
		next.ServeHTTP(w, r)
	})
}

func TestMiddlewareFunc(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("Index")
		w.Write([]byte("Hello world"))
	})

	authMiddleware := AuthMiddleware(router)
	errMiddleware := ErrMiddleware(authMiddleware)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	errMiddleware.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Hello world", string(body))
}

// Middleware sebagai objek
type MiddlewareStruct struct {
	Handler http.Handler
}
func newMiddleware(h http.Handler) *MiddlewareStruct {
	return &MiddlewareStruct{Handler: h}
}
func (middleware *MiddlewareStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Middleware 1: Ini perintah sebelum diteruskan handler selanjutnya")

	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddlewareObj(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("Hello world"))
	})

	middleware := newMiddleware(router)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	middleware.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Hello world", string(body))
}
