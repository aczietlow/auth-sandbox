package clientServer

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	// compile templates
	t := compileTemplates()
	e := echo.New()
	e.Static("/css", "web/css")
	e.Static("/js", "web/js")
	e.Static("/assets", "web/assets")
	e.Renderer = t

	// register routes
	e.GET("/", index)
	e.GET("/oauth2/callback", oauthCallback)

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// start server
	e.Logger.Fatal(e.Start(":1322"))
}

func index(c echo.Context) error {
	// redirect := ""
	return c.Render(http.StatusOK, "index", nil)
}

func oauthCallback(c echo.Context) error {
	return c.String(http.StatusOK, "ain't no hollar back")
}
