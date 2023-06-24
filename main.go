package main

import (
	"flag"
	"github.com/gookit/slog"
	"table-export/config"
	"table-export/consts"
	"table-export/export"
	"table-export/meta"
)

func main() {

	slog.SetLogLevel(slog.InfoLevel)

	flag.Parse()

	if params.help {
		usage()
		return
	}

	if params.version {
		consts.PrintBuild()
	}

	config.ParseConfig(params.confFile)

	if params.genSource != "" {
		genMeta := meta.NewGenMeta(params.genSource)
		genMeta.Run()
	}

	if params.mode != "" {
		entry := export.NewEntry(params.mode, params.extra)
		entry.Run()
	}

}
