package web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// data
type LayoutPage struct{
	Title string
	User map[string]interface{}
}

var layoutPage LayoutPage = LayoutPage{
	Title: "Layout Page",
	User : map[string]interface{}{
			"Name": "Helmi",
			"Age": 23,
	},
}

// embed
//go:embed templates/*.html
var layoutTemplates embed.FS

// handler
func layoutHandler(w http.ResponseWriter, r *http.Request) {
	// create template
	tmpl := template.Must(template.ParseFS(layoutTemplates, "templates/*.html"))

	// data
	data := &layoutPage

	// execute templates and send to response
	tmpl.ExecuteTemplate(w, "layout", *data)
}

func TestLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:8080", nil)
	recorder := httptest.NewRecorder()
	
	layoutHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}


/* func TestLayout(t *testing.T) {
	http.HandleFunc("/", layoutHandler)

	addr := "localhost:8080"
	fmt.Println("running at: ", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
} */