package openidconnect

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Write the configuration to the file
	text := []byte("redirectURL: http://localhost:9001/redirect\nappURL: http://localhost:9001/\nid: openid\nauthURL: http://localhost/auth")
	if _, err := tmpfile.Write(text); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Load the configuration
	conf, err := loadConfig(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Check the configuration values
	if conf.RedirectURL != "http://localhost:9001/redirect" {
		t.Errorf("expected redirectURL to be http://localhost:9001/redirect, got %s", conf.URL)
	}
	if conf.URL != "http://localhost:9001/" {
		t.Errorf("expected URL to be http://localhost:9001/, got %s", conf.URL)
	}
	if conf.ID != "openid" {
		t.Errorf("expected id to be openid, got %s", conf.ID)
	}
	if conf.AuthURL != "http://localhost/auth" {
		t.Errorf("expected authURL to be http://localhost/auth, got %s", conf.AuthURL)
	}
	if conf.Secret != "pass" {
		t.Errorf("expected authURL to be pass, got %s", conf.Secret)
	}
}
