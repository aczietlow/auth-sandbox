package handler

import (
	"fmt"
	"net/http"
	"net/url"

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
		InfoDump:    prettyPrintUrl(authUrl),
	})
}

func prettyPrintUrl(raw string) string {
	u, _ := url.Parse(raw)
	queryParams := u.Query()
	u.RawQuery = ""

	infoURL := fmt.Sprintf("%s \n\n", u.String())
	for k, v := range queryParams {
		infoURL += fmt.Sprintf("%s\t\t\t = %s\n", k, v)
	}

	return infoURL
}
