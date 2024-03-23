package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/helmimuzkr/templ/handler"
	"github.com/helmimuzkr/templ/utill"
)

//go:embed view/*.html
var viewFiles embed.FS

func main() {
	mux := http.NewServeMux()

	utill.InitPostgresDB()

	templEngine := utill.NewTemplateEngine(viewFiles, "view/*.html")
	templEngine.GenerateTemplate()

	mux.HandleFunc("GET /{params}", handler.LogMiddleware(handler.HelloWorldHandler))

	mux.HandleFunc("GET /users", handler.LogMiddleware(handler.UserInformationHandler))
	mux.HandleFunc("GET /users/list", handler.LogMiddleware(handler.UserListHandler))
	mux.HandleFunc("POST /users", handler.LogMiddleware(handler.RegisterUserHandler))

	port := ":6969"
	fmt.Printf("listening at port: %s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
