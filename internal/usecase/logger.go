package usecase

import (
	"github.com/gettmure/go-click/internal/config"
	"github.com/gettmure/go-click/internal/lib"
	"github.com/gettmure/go-click/internal/logger"
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
	u.log = u.log.With("operation", op)

	lib.LogClickSuccess(u.log, body, u.cfg.LogBody)
}
