package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// -----------
// TEMPLATE
// -----------
// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
	debug     bool
	location  string
}

// Render renders a template document
func NewRenderer(location string, debug bool) *TemplateRenderer {
	tpl := new(TemplateRenderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}
func (t *TemplateRenderer) ReloadTemplates() {
	t.templates = template.Must(template.ParseGlob(t.location))
}
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func templateHandler(c echo.Context) error {
	// Buat data yang akan di passing ke image.html
	data := make(map[string]interface{})
	data["filename"] = "avatar.jpg"

	return c.Render(http.StatusOK, "image.html", data)
}

// -----------
// UPLOAD
// -----------
func uploadHandler(c echo.Context) error {
	//-----------
	// Read from field
	//-----------
	// caption := c.FormValue("caption")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return nil
	}
	defer src.Close()

	path := "./assets/image/"
	filename := path + file.Filename

	// Destination
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy source(from http) to destination(local or server)
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, fmt.Sprintln("file success ", filename))
}

// -----------
// STATIC REDIRECT FILE IMAGE
// -----------
func profileHandler(c echo.Context) error {
	filename := c.QueryParam("name")

	// file, err := os.Open("./assets/image/" + filename)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	path := "./assets/image/" + filename

	return c.File(path)
}

func main() {

	// New echo
	e := echo.New()

	// Belajar upload file
	e.POST("/upload", uploadHandler)

	// Belajar echo html template
	e.Renderer = NewRenderer("./public/*.html", true)
	e.GET("/image", templateHandler)

	// Belajar  static
	e.Static("/", "public")

	// Belajar static file
	e.File("/avatar", "assets/image/avatar.jpg")

	// Belajar static file dengan handler
	e.GET("/users/", profileHandler)

	// server start
	if err := e.Start(":8000"); err != nil {
		fmt.Println(err)
		return
	}

}

// dsn := "root:helmi@tcp(localhost:3306)/pagination?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	log.Println(err)
// }
