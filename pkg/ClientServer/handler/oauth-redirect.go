package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) OAuthCallback(c echo.Context) error {
	return c.String(http.StatusOK, "ain't no hollar back")
}
