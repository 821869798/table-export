package util

import (
	"github.com/gookit/slog"
	"time"
)

func TimeCost(start time.Time, formatStr string) {
	tc := time.Since(start)
	slog.Infof(formatStr, tc)
}
