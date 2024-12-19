package main

import (
	configuration "secret-satana/configs"
	"secret-satana/database"
	"secret-satana/routes"
	satana_selection "secret-satana/satana-selection-logic"

	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom HTML template renderer for Echo
type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Initialize the database
	database.InitDatabase()

	if configuration.ReselectSatana {
		satana_selection.ReselectSatanas()
	}

	// Initialize Echo
	e := echo.New()

	// Set up template renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/public", "templates/public")

	// Routes
	e.Use(routes.JwtMiddleware)

	routes.RegisterLoginRoutes(e)
	routes.RegisterIndexRoutes(e)
	routes.RegisterParticipateRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
