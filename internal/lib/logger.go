package lib

import (
	"github.com/gettmure/go-click/internal/logger"
	"golang.org/x/exp/slog"
)

func LogClickSuccess(log logger.Logger, body []byte, logBody bool) {
	if logBody {
		bodyString := string(body[:])
		log.Info(
			"click success",
			slog.String("status", "ok"),
			slog.String("body", bodyString),
		)

		return
	}

	log.Info(
		"click success",
		slog.String("status", "ok"),
	)
}

func LogClickError(log logger.Logger, err error) {
	log.Info(
		"click error",
		slog.String("status", "error"),
		slog.String("description", err.Error()),
	)
}
