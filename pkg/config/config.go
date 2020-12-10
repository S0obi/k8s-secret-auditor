package config

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"
)

// Config : Global Config
type Config struct {
	IgnoredNamespaces []string `yaml:"ignoredNamespaces"`
	PasswordPatterns  []string `yaml:"passwordPatterns"`
	Policy            struct {
		Length  int     `yaml:"length"`
		Entropy float64 `yaml:"entropy"`
	}
}

// NewConfig : Load current config (config.yaml)
func NewConfig(configPath string) *Config {
	config := new(Config)

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err.Error())
	}

	if err = yaml.UnmarshalStrict(data, &config); err != nil {
		panic(err.Error())
	}

	if len(config.PasswordPatterns) == 0 {
		config.PasswordPatterns = append(config.PasswordPatterns, "password", "pass", "pwd")
	}

	return config
}
