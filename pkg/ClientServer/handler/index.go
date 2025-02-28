package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Index(c echo.Context) error {
	data := h.OIDC.GetAuthorizationServerURL()
	return c.Render(http.StatusOK, "index", data)
}
