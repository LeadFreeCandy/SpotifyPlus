package internal

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	clientID    string `yaml:"clientID"`
	redirectURI string `yaml:"redirectURI"`
	serverPort  int16  `yaml:"port"`
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
