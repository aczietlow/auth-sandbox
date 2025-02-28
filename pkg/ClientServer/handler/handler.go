package handler

import openidconnect "github.com/aczietlow/auth-sandbox/pkg/OpenIDConnect"

type Handler struct {
	// This needs to be new per user session b/c of the state parameter
	OIDC *openidconnect.OIDC
}
