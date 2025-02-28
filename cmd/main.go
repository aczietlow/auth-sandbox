package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"

	openidconnect "github.com/aczietlow/auth-sandbox/pkg/OpenIDConnect"
	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func compileTemplates() *Templates {
	templateDirs := []string{
		"web/views",
		"web/views/partials",
	}

	templates := template.New("")
	for _, dir := range templateDirs {
		// Using filepath.Glob to get all template files
		files, err := filepath.Glob(filepath.Join(dir, "*.html"))
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			log.Printf("Adding %v to temaplte path", file)
			// Parse each template file and add to the template set
			_, err = templates.ParseFiles(file)
			if err != nil {
				panic(err)
			}
		}
	}

	return &Templates{
		templates: templates,
	}
}

func main() {
	// compile templates
	t := compileTemplates()
	e := echo.New()
	e.Static("/css", "web/css")
	e.Static("/js", "web/js")
	e.Static("/assets", "web/assets")
	e.Renderer = t

	// register routes
	e.GET("/", index)

	// Get fancy client
	oidc, err := openidconnect.New("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("New OIDC Client created: %s", oidc.ClientID)

	// start server
	e.Logger.Fatal(e.Start(":1322"))
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
