package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config of service
type Config struct {
	PostgresURL string `json:"postgresURL"`
	Host        string `json:"host"`
	Port        string `json:"port"`
}

// ParseConfig of service
func ParseConfig(configPath string) (*Config, error) {
	fileBody, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(fileBody, &cfg)
	return &cfg, nil
}
