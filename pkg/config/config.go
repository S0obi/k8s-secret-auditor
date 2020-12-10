package config

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"
)

// Config : Global Config
type Config struct {
	Policy struct {
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

	return config
}
