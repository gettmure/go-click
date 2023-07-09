package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	SiteConfig    SiteConfig    `yaml:"site"`
	RuntimeConfig RuntimeConfig `yaml:"runtime"`
}

type SiteConfig struct {
	Url string `yaml:"url"`
}

type RuntimeConfig struct {
	Env     string `yaml:"env"`
	LogBody bool   `yaml:"logBody"`
}

func MustLoad() Config {
	configPath := flag.String("config", "internal/config/config.yaml", "path to config.yaml file")
	flag.Parse()

	_, err := os.Stat(*configPath)
	if os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", *configPath)
	}

	cfg := &Config{}

	err = cleanenv.ReadConfig(*configPath, cfg)
	if err != nil {
		log.Fatalf("failed to load config file: %s", err)
	}

	return *cfg
}
