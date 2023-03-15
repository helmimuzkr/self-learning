package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testing"
)

// kalau mau declare struct pointer, harus pake new(Struct) supaya data pada struct tidak nil.
// jika nil, akan menyebabkan panic error
// contoh
// var products *Products // NIL, and will panic error
// data di dalam struct sama seperti deklarasi struct non pointer
var products = new(Products) // NOT NIL
var cookie = new(http.Cookie)

var usernameValid string = "helmi"
var isLogin bool = false

// Handler
func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("halaman login"))
}
func loginProcess(w http.ResponseWriter, r *http.Request) {	
	if r.PostFormValue("username") != usernameValid {
		fmt.Println("username salah")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	// set cookie
	cookie.Name = "username"
	cookie.Value = r.PostFormValue("username")
	cookie.Path = "/"
	cookie.MaxAge = 0
	isLogin = true 
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
}
func logoutProcess(w http.ResponseWriter, r *http.Request) {
	// set cookie
	cookie.Name = "username"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.MaxAge = -1
	isLogin = false
	http.SetCookie(w, cookie)
	
	fmt.Println("logout completed, redirect to index")
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func product(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halaman Product\n"))

	for _, v := range products.Products {
		productStr := strconv.Itoa(v.ID) + " " + v.Name + " " + strconv.Itoa(v.Price) + "\n"

		w.Write([]byte(productStr))
	}
}
func addProduct(w http.ResponseWriter, r *http.Request) {
	// Get value form post form
	id, _ := strconv.Atoi(r.PostFormValue("product_id"))
	name := r.PostFormValue("product_name")
	price,_  := strconv.Atoi(r.PostFormValue("product_price"))

	// Membuat struct product dan assign form post value ke field
	product := &Product{
		ID: id,
		Name: name,
		Price: price,
	}

	// Pass product ke method struct Products
	products.addProduct(product)
	
	http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
}

// Middleware
func authUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.RequestURI()
		fmt.Println(url)

		if (url == "/" || url == "/login")  {
			if isLogin {
				http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
				return
			}
			
			next.ServeHTTP(w, r)
			return
		}

		if url == "/logout" {
			next.ServeHTTP(w, r)
			return
		}

		if !isLogin {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func TestMiddleware(t *testing.T) {	
	mux := http.NewServeMux()

	mux.HandleFunc("/", login)
	mux.HandleFunc("/login", loginProcess)
	mux.HandleFunc("/logout", logoutProcess)
	mux.HandleFunc("/product", product)
	mux.HandleFunc("/tambah", addProduct)

	handler := http.Handler(mux)
	handler = authUserMiddleware(handler)

	// Setup server
	port := ":8080"
	fmt.Println("Server listening on port ", port)
	if err := http.ListenAndServe("localhost"+port, handler); err != nil {log.Fatal(err)}
}

