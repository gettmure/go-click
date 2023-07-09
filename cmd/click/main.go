package main

import (
	lg "log"

	"github.com/gettmure/go-click/internal/clicker"
	"github.com/gettmure/go-click/internal/config"
	"github.com/gettmure/go-click/internal/logger"
	"golang.org/x/exp/slog"
)

func main() {
	const op = "cmd.click.main"
	cfg := config.MustLoad()

	log, err := logger.New(cfg.RuntimeConfig.Env)
	if err != nil {
		lg.Fatalf("failed to init logger: %s", err)
	}
	log = log.With("operation", op)
	log.Debug("debug messages are enabled")

	log.Info(
		"config loaded",
		slog.String("env", cfg.RuntimeConfig.Env),
		slog.String("site_url", cfg.SiteConfig.Url),
	)

	clicker := clicker.New()
	body, err := clicker.Click(cfg.SiteConfig.Url)
	if err != nil {
		log.Info(
			"click error",
			slog.String("status", "error"),
			slog.String("description", err.Error()),
		)
	}

	if cfg.RuntimeConfig.LogBody {
		bodyString := string(body[:])
		log.Info(
			"click success",
			slog.String("status", "ok"),
			slog.String("body", bodyString),
		)
	} else {
		log.Info(
			"click success",
			slog.String("status", "ok"),
		)
	}
}
