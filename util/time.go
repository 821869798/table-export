package util

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func TimeCost(start time.Time, formatStr string) {
	tc := time.Since(start)
	log.Infof(formatStr, tc)
}
