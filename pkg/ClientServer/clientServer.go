package clientServer

import (
	"github.com/aczietlow/auth-sandbox/pkg/ClientServer/handler"
	openidconnect "github.com/aczietlow/auth-sandbox/pkg/OpenIDConnect"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(c openidconnect.OIDC) {
	// compile templates
	t := compileTemplates()
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.Static("/css", "web/css")
	e.Static("/js", "web/js")
	e.Static("/assets", "web/assets")
	e.Renderer = t

	h := &handler.Handler{
		OIDC: &c,
	}

	// register routes
	e.GET("/", h.Index)
	e.GET("/oauth2/callback", h.OAuthCallback)
	e.GET("/authflow", h.AuthFlow)

	// start server
	e.Logger.Fatal(e.Start(":1322"))
}
