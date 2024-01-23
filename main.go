package main

import (
	"flag"
	"github.com/821869798/fankit/fanpath"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/constant"
	"github.com/821869798/table-export/export"
	"github.com/821869798/table-export/meta"
	"github.com/gookit/slog"
)

func main() {

	slog.SetLogLevel(slog.InfoLevel)

	err := fanpath.InitExecutePath()
	if err != nil {
		slog.Fatal(err)
	}

	flag.Parse()

	if params.help {
		usage()
		return
	}

	if params.version {
		constant.PrintBuild()
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
