package renderer

import (
	"io"
	"text/template"

	echo "github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom renderer for Echo that uses the Go template engine.
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template with the given name and data.
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// New returns a default new template renderer
func New() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}