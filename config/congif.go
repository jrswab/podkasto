package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	ApiKey    string
	ApiSecret string
}

func (c *Config) Load() {
	var (
		err         error
		configBytes []byte
	)

	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir != "" {
		configBytes, err = os.ReadFile(fmt.Sprintf("%s/podkasto.conf", configDir))
		if err != nil {
			log.Fatalf("could not load $XDG_CONFIG_HOME/podkasto.conf: %s", err)
		}
	}

	if configDir == "" {
		configDir = os.Getenv("HOME")

		configBytes, err = os.ReadFile(fmt.Sprintf("%s/.podkasto.conf", configDir))
		if err != nil {
			log.Fatalf("could not load $HOME/.podkasto.conf: %s", err)
		}
	}

	configSettings := strings.Split(string(configBytes), "\n")

	for _, line := range configSettings {
		lineData := strings.Split(line, " ")

		if len(lineData) <= 1 {
			continue
		}

		switch lineData[0] {
		case "apiKey":
			c.ApiKey = lineData[1]
		case "apiSecret":
			c.ApiSecret = lineData[1]

		}
	}
}
