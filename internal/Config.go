package internal

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	ClientID    string `yaml:"clientID"`
	RedirectURI string `yaml:"redirectURI"`
	ServerPort  int16  `yaml:"port"`
}

func ParseYamlConfig(cfg *Config, configPath string) error {
	f, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	return err
}
