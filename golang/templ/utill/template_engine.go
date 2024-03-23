package utill

import (
	"embed"
	"html/template"
	"net/http"
)

var TemplEngine *TemplateEngine

type TemplateEngine struct {
	templ *template.Template

	fs      embed.FS
	pattern string
}

func NewTemplateEngine(templateFiles embed.FS, pattern string) *TemplateEngine {
	return &TemplateEngine{
		fs:      templateFiles,
		pattern: pattern,
	}
}

func (t *TemplateEngine) GenerateTemplate() {
	if TemplEngine != nil {
		return
	}

	t.templ = template.Must(template.ParseFS(t.fs, t.pattern))

	TemplEngine = t
}

func (t *TemplateEngine) ExecuteTemplate(w http.ResponseWriter, name string, data any) error {
	return t.templ.ExecuteTemplate(w, name, data)
}
