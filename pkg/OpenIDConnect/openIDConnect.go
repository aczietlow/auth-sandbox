package openidconnect

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
)

type OIDC struct {
	ClientID    string
	ClientURL   string
	AuthURL     string
	RedirectURL string
	Secret      string
	State       string
	// Verifier    string
	// Token       Token
}

func New(filepath string) (OIDC, error) {
	c, err := loadConfig(filepath)
	if err != nil {
		return OIDC{}, err
	}

	state, err := generateState(10)
	if err != nil {
		return OIDC{}, err
	}

	o := OIDC{
		ClientID:    c.ID,
		ClientURL:   c.URL,
		AuthURL:     c.AuthURL,
		RedirectURL: c.RedirectURL,
		Secret:      c.Secret,
		State:       state,
	}

	return o, nil
}

func generateState(n int) (string, error) {
	data := make([]byte, n)
	if _, err := rand.Read(data); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func generateCodeVerifier() (string, error) {
	data := make([]byte, 32)
	if _, err := rand.Read(data); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(data), nil
}

func s256CodeChallenge(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	return base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
}

func (o *OIDC) GetAuthorizationServerURL() string {
	// verifier, _ := generateCodeVerifier()
	//
	// c.Verifier = verifier
	params := o.getHandshakeParams()
	u, err := url.Parse(o.AuthURL)
	if err != nil {
		// TODO: do something
		fmt.Println(err.Error())
	}
	u.RawQuery = params.Encode()

	return u.String()
}

func (o *OIDC) getHandshakeParams() url.Values {
	// BASE64URL-ENCODE(SHA256(ASCII(code_verifier)))
	// codeChallenge := s256CodeChallenge(c.Verifier)

	request := url.Values{
		"response_type": []string{"code"},
		"client_id":     []string{o.ClientID},
		"redirect_uri":  []string{o.RedirectURL},
		// TODO: Make scopes configurable
		"scope": []string{"profile email"},
		"state": []string{o.State},
		// "code_challenge":        []string{codeChallenge},
		// "code_challenge_method": []string{"S256"},
	}

	return request
}
