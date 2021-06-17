package cli

import (
	"os"

	"github.com/BurntSushi/toml"
)

func ReadConfig(path string) *ClientConfig {
	var config ClientConfig

	if _, err := toml.DecodeFile(path, &config); err != nil {
		PrintServerError("Invalid client configuration file provided.")
		os.Exit(1)
	}

	return &config
}

type ClientConfig struct {
	BgColor   string `json:"bg-color"`
	Count     int    `json:"count"`
	Duration  int    `json:"duration"`
	FontColor string `json:"font-color"`
	FontSize  int    `json:"font-size"`
}
