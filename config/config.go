package config

import (
	_ "embed"
	"fmt"

	"github.com/go-yaml/yaml"
)

//go:embed config.yml
var configInput string

type ExternalAPI int
type Configuration interface {
	GetAPIKey(target ExternalAPI) string
}

type ConfigData struct {
	ApiKey APIKeys `yaml:"API_KEY"`
}
type APIKeys struct {
	DeepL     string `yaml:"DEEPL"`
	APINinjas string `yaml:"API_NINJAS"`
}

const (
	DEEPL ExternalAPI = iota
	API_NINJAS
)

func GetConfiguration() Configuration {
	var conf ConfigData
	err := yaml.Unmarshal([]byte(configInput), &conf)
	if err != nil {
		fmt.Println("failed parse configuration YAML. Please check ./config/config.yaml")
	}
	return conf
}

func (c ConfigData) GetAPIKey(target ExternalAPI) string {
	if target == DEEPL {
		return c.ApiKey.DeepL
	} else if target == API_NINJAS {
		return c.ApiKey.APINinjas
	}
	return ""
}
