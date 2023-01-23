package web

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"testing"
)

// Contoh data yang akan dikirim ke html
type Data struct {
	Title string
	Content map[string]interface{}
}

// Test Template
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Method template.ParseGlob() dipanggil dengan parameter yaitu pattern path "views/*". 
	// Method ini digunakan untuk memparsing semua file yang match dengan pattern yang ditentukan, 
	// dan fungsi ini mengembalikan 2 objek: *template.Template & error.
	// Pattern path pada fungsi template.ParseGlob() nantinya akan di proses oleh filepath.Glob()
	tmpl, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		panic(err)
	}

	// Method template.ParseFiles() dipanggil dengan parameter yaitu pattern "directory/index.html"
	// Method ini digunakan untuk melukan parsing suatu file dan akan mengembalikan 2 objek: *template.Template & error.
	// tmpl, err := template.ParseFiles("./templates/index.html")
	// if err != nil {
	// 	panic(err)
	// }

	// Menggunakan method Must() agar tidak menghandle error manual
	// t := template.Must(template.ParseGlob("./templates/*.html"))



	data := Data{
		Title: "Belajar templating golang",
		Content: map[string]interface{}{
			"Name": "Helmi",
			"Umur": 23,
		},
	}

	// Nanti di file htmlnya bisa diakses dengan template test => {{.NamaField}}

	tmpl.ExecuteTemplate(w, "index.html", data)
}
func TestTemplate(t *testing.T) {
	http.HandleFunc("/", indexHandler)

	addr := "localhost:8080"
	fmt.Println("running server at: ", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}


// Test Template With Embed File
// Kenapa harus menggunakan embed? agar saat kita selesai compile aplikasinya
// kita jadi tidak harus bergantung pada file eksternal.
// Karena saat dilakukan proses embed templates, 
// file golagnya disempen dalam variable templates yang bertipe data embed.FS

//go:embed templates/*.html
var templatesWithEmbed embed.FS

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// create template
	tmpl := template.Must(template.ParseFS(templatesWithEmbed, "templates/*.html"))

	// Data
	data := Data{
		Title: "Template dengan Embed",
		Content: map[string]interface{}{
			"Name": "Helmi",
		},
	}

	// execute template
	tmpl.ExecuteTemplate(w, "index.html", data)
}
func TestTemplateEmbed(t *testing.T) {
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}