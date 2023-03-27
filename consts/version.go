package consts

import (
	log "github.com/sirupsen/logrus"
)

var (
	Version = "0.0.1"
)

func PrintBuild() {
	log.WithFields(log.Fields{
		"Version": Version,
	}).Info()
}
