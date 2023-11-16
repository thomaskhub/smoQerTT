package utils

import (
	"flag"
	"log"
	"os"

	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"gopkg.in/yaml.v2"
)

type Ws struct {
	EnableTLS bool   `yaml:"enable_tls"`
	WsKey     string `yaml:"ws_key"`
	WsCrt     string `yaml:"ws_crt"`
	Listening string `yaml:"listening"`
}

type Tls struct {
	EnableTLS bool   `yaml:"enable_tls"`
	TlsKey    string `yaml:"tls_key"`
	TlsCrt    string `yaml:"tls_crt"`
	Port      int    `yaml:"port"`
	Listening string `yaml:"listening"`
}

type Config struct {
	Tls    Tls         `yaml:"tls"`
	Ws     Ws          `yaml:"ws"`
	Ledger auth.Ledger `yaml:"ledger"`
}

func ParseConfig(file string) *Config {
	// Read the yaml file
	cfg := Config{}
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Unmarshal the yaml file into the config struct
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Parse the flags
	flag.Parse()

	return &cfg
}
