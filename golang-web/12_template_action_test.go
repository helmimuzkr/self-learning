package web

import (
	"embed"
	"fmt"
	"net/http"
	"testing"
	"text/template"
)

type People struct {
	Name string
	Hobbies []string
}

type ActionPage struct {
	Title string
	People People
}

//go:embed templates/*.html
var actionTemplates embed.FS
// pattern for parsing
var patternTemplates string = "templates/*.html"

// data
var data ActionPage = ActionPage{
	Title: "Belajar Action: If",
	People: People{
		Name: "Helmi",
		Hobbies: []string{
			"Playing Games", "Swimming", "Sleeping",
		},
	},
}


// If Action
func ifHandler(writer http.ResponseWriter, request *http.Request) {
	// parsing template
	tmpl := template.Must(template.ParseFS(actionTemplates, "templates/*.html"))

	// data
	data := &data

	// execution and send to response
	tmpl.ExecuteTemplate(writer, "action-if.html", *data)
}

// Range Action
func rangeHandler(writer http.ResponseWriter, request *http.Request) {
	// parsing template
	tmpl := template.Must(template.ParseFS(actionTemplates, patternTemplates))

	// data
	// Untuk template Range, data yang diiterasi bisa dalam bentuk
	// struct, slices, map, array, maupun channel
	data := &data

	// exec and send to response
	tmpl.ExecuteTemplate(writer, "action-range.html", *data)
}	

// Compare Action
func compareHandler(w http.ResponseWriter, r *http.Request) {
	// parsing template
	tmpl := template.Must(template.ParseFS(actionTemplates, patternTemplates))

	// data
	data := &data

	// exec and send to response
	tmpl.ExecuteTemplate(w, "action-compare.html", data)
}

// With action
func withHandler(w http.ResponseWriter, r *http.Request) {
	// create template
	tmpl := template.Must(template.ParseFS(actionTemplates, patternTemplates))

	// data
	data := &data

	// exec templates and send to response
	tmpl.ExecuteTemplate(w, "action-with.html", data)
}


// Test the Action
func TestAction(t *testing.T) {
	// create route
	http.HandleFunc("/if", ifHandler) 
	http.HandleFunc("/range", rangeHandler)
	http.HandleFunc("/compare", compareHandler)
	http.HandleFunc("/with", withHandler)

	// address
	addr := "localhost:8080"
	fmt.Println("Running server at: ", addr)

	// running server
	http.ListenAndServe(addr, nil)
}