package logger

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

func loadDev() (*slog.Logger, error) {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	), nil
}

func loadProd() (*slog.Logger, error) {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create logs directories for path %s: %s", dirPath, err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %s", err)
	}
	w := bufio.NewWriter(f)

	return slog.New(
		slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelDebug}),
	), nil
}
