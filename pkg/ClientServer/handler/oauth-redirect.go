package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type data struct {
	InfoDump        string
	ExchangeCodeUrl string
}

func (h *Handler) OAuthCallback(c echo.Context) error {
	url := c.Request().URL.String()
	d := h.OIDC.State
	return c.Render(http.StatusOK, "oidc-callback", data{
		InfoDump:        h.PrettyPrintUrl(url) + "\n" + d,
		ExchangeCodeUrl: "",
	})
}
