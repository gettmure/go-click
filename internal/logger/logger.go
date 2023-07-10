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

type Logger = *slog.Logger

func New(env string) (Logger, error) {
	switch env {
	case envDev:
		return loadDev()
	case envProd:
		return loadProd()
	default:
		return loadDev()
	}
}
