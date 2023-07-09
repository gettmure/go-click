package logger

import (
	"golang.org/x/exp/slog"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

const (
	dirPath  = "./internal/logs"
	filePath = "./internal/logs/log.prod"
)

func New(env string) (*slog.Logger, error) {
	switch env {
	case envDev:
		return loadDev()
	case envProd:
		return loadProd()
	default:
		return loadDev()
	}
}
