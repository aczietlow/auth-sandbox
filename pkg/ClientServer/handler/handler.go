package handler

import (
	"fmt"
	"net/url"

	openidconnect "github.com/aczietlow/auth-sandbox/pkg/OpenIDConnect"
)

type Handler struct {
	// This needs to be new per user session b/c of the state parameter
	OIDC *openidconnect.OIDC
}

func (h *Handler) PrettyPrintUrl(raw string) string {
	u, _ := url.Parse(raw)
	queryParams := u.Query()
	u.RawQuery = ""

	infoURL := fmt.Sprintf("%s \n\n", u.String())
	for k, v := range queryParams {
		infoURL += fmt.Sprintf("%s\t\t\t = %s\n", k, v)
	}

	return infoURL
}
