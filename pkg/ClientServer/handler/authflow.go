package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type redirectData struct {
	RedirectUri string
	InfoDump    string
}

func (h *Handler) AuthFlow(c echo.Context) error {
	authUrl := h.OIDC.GetAuthorizationServerURL()
	return c.Render(http.StatusOK, "authflow", redirectData{
		RedirectUri: authUrl,
		InfoDump:    h.PrettyPrintUrl(authUrl),
	})
}
