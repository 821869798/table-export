package constant

import (
	"github.com/gookit/slog"
)

var (
	Version = "1.1.0"
)

func PrintBuild() {
	slog.Infof("Version:%s", Version)
}
