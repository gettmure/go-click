package main

import (
	"github.com/gettmure/go-click/internal/config"
	"github.com/gettmure/go-click/internal/logger"
	"github.com/gettmure/go-click/internal/usecase"
	"github.com/gettmure/go-click/pkg/clicker"
	"golang.org/x/exp/slog"
)

func main() {
	const op = "cmd.clicker.main"
	cfg := config.MustLoad()

	log, err := logger.New(cfg.RuntimeConfig.Env)
	if err != nil {
		panic(err)
	}
	log = log.With("operation", op)
	log.Debug("debug messages are enabled")

	log.Info(
		"config loaded",
		slog.String("env", cfg.RuntimeConfig.Env),
		slog.String("site_url", cfg.SiteConfig.Url),
	)

	clicker, err := clicker.New()
	if err != nil {
		panic(err)
	}

	clickerUsecase := usecase.NewClickerUsecase(clicker, log, cfg.SiteConfig)
	loggerUsecase := usecase.NewLoggerUsecase(log, cfg.LoggerConfig)

	body := clickerUsecase.Click()
	loggerUsecase.LogBody(body)
}
