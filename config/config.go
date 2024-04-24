package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Url        string   `json:"url"`
	Priorities []string `json:"priorities"`
	Exclude    []string `json:"exclude"`
	Limit      int      `json:"limit"`
}

func ReadConfig() Config {
	file, err := os.Open("cozy-cache.json")
	if err != nil {
		panic("Config file not found")
	}

	parser := json.NewDecoder(file)
	config := Config{}

	if err = parser.Decode(&config); err != nil {
		panic("Unable to decode json file")
	}

	return config
}
