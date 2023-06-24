package consts

import (
	"github.com/gookit/slog"
)

var (
	Version = "0.0.1"
)

func PrintBuild() {
	slog.Infof("Version:%s", Version)
}
