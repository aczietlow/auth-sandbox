package openidconnect

type OIDC struct {
	ClientID    string
	ClientURL   string
	AuthURL     string
	RedirectURL string
	Secret      string
	// State       string
	// Verifier    string
	// Token       Token
}

func New(filepath string) (OIDC, error) {
	c, err := loadConfig(filepath)
	if err != nil {
		return OIDC{}, err
	}

	o := OIDC{
		ClientID:    c.ID,
		ClientURL:   c.URL,
		AuthURL:     c.AuthURL,
		RedirectURL: c.RedirectURL,
		Secret:      c.Secret,
	}

	return o, nil
}
