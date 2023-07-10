package usecase

import (
	"github.com/gettmure/go-click/internal/clicker"
	"github.com/gettmure/go-click/internal/config"
	"github.com/gettmure/go-click/internal/logger"
	"golang.org/x/exp/slog"
)

type ClickerUsecase interface {
	FetchBody() []byte
}

type clickerUsecase struct {
	clicker clicker.Clicker
	log     logger.Logger
	cfg     config.SiteConfig
}

func NewClickerUsecase(clicker clicker.Clicker, log logger.Logger, cfg config.SiteConfig) ClickerUsecase {
	return &clickerUsecase{
		clicker: clicker,
		log:     log,
		cfg:     cfg,
	}
}

func (u *clickerUsecase) FetchBody() []byte {
	var body []byte

	body, err := u.clicker.Click(u.cfg.Url)
	if err != nil {
		u.log.Info(
			"click error",
			slog.String("status", "error"),
			slog.String("description", err.Error()),
		)

		return body
	}

	return body
}
