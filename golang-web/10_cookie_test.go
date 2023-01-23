package web

import (
	"fmt"
	"net/http"
	"testing"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-MZKR-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Suskses membuat Cookie")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _:= r.Cookie("name")
	value := cookie.Value
	fmt.Fprint(w, value)
}

func TestCookie(t *testing.T){
	http.HandleFunc("/get", getCookie)
	http.HandleFunc("/set", setCookie)

	addr := "localhost:8080"
	fmt.Println("menjalankan server: ", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
