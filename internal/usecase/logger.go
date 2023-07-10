package usecase

import (
	"github.com/gettmure/go-click/internal/config"
	"github.com/gettmure/go-click/internal/logger"
	"golang.org/x/exp/slog"
)

type LoggerUsecase interface {
	LogBody(body []byte)
}

type loggerUsecase struct {
	log logger.Logger
	cfg config.LoggerConfig
}

func NewLoggerUsecase(log logger.Logger, cfg config.LoggerConfig) LoggerUsecase {
	return &loggerUsecase{
		log: log,
		cfg: cfg,
	}
}

func (u *loggerUsecase) LogBody(body []byte) {
	const op = "internal.usecase.logger.LogBody"
	u.log = u.log.With(op)

	if u.cfg.LogBody {
		bodyString := string(body[:])
		u.log.Info(
			"click success",
			slog.String("status", "ok"),
			slog.String("body", bodyString),
		)
	} else {
		u.log.Info(
			"click success",
			slog.String("status", "ok"),
		)
	}
}
