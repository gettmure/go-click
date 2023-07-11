package usecase

import (
	"math/rand"
	"time"

	"github.com/gettmure/go-click/internal/config"
	"github.com/gettmure/go-click/internal/lib"
	"github.com/gettmure/go-click/internal/logger"
	"github.com/gettmure/go-click/pkg/clicker"
)

type ClickerUsecase interface {
	Click() []byte
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

func (u *clickerUsecase) Click() []byte {
	const op = "internal.usecase.clicker.Click"
	u.log.With("operation", op)

	body, err := u.clicker.Click(u.cfg.Url)
	if err != nil {
		lib.LogClickError(u.log, err)

		return []byte{}
	}

	if u.cfg.ExtraLinks.Enabled {
		clickExtraLinks(u)
	}

	return body
}

func clickExtraLinks(u *clickerUsecase) {
	const op = "internal.usecase.clicker.clickExtraLinks"
	u.log.With("operation", op)

	links := u.cfg.ExtraLinks.Links

	if u.cfg.ExtraLinks.Enabled {
		if u.cfg.ExtraLinks.Random {
			src := rand.NewSource(time.Now().Unix())
			rand := rand.New(src)
			link := links[rand.Intn(len(links))]

			_, err := u.clicker.Click(link)
			if err != nil {
				lib.LogClickError(u.log, err)
			}

			return
		}

		for _, link := range u.cfg.ExtraLinks.Links {
			_, err := u.clicker.Click(link)
			if err != nil {
				lib.LogClickError(u.log, err)
			}
		}
	}
}
