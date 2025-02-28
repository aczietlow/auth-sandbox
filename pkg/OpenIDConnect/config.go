package openidconnect

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ID          string `yaml:"id"`
	RedirectURL string `yaml:"redirectURL"`
	URL         string `yaml:"appURL"`
	AuthURL     string `yaml:"authURL"`
	Secret      string `yaml:"secret"`
}

func loadConfig(filepath string) (*Config, error) {
	configFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	c := Config{}
	if err = yaml.Unmarshal(configFile, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
