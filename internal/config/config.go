package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	SiteConfig    SiteConfig    `yaml:"site"`
	RuntimeConfig RuntimeConfig `yaml:"runtime"`
	LoggerConfig  LoggerConfig  `yaml:"logger"`
}

type SiteConfig struct {
	Url string `yaml:"url"`
}

type RuntimeConfig struct {
	Env string `yaml:"env"`
}

type LoggerConfig struct {
	LogBody bool `yaml:"logBody"`
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_FILE")
	if len(configPath) == 0 {
		panic("empty config path")
	}

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		panic(err)
	}

	cfg := &Config{}

	err = cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		panic(err)
	}

	return *cfg
}
