package main

import (
	"flag"
	"github.com/shiena/ansicolor"
	log "github.com/sirupsen/logrus"
	"os"
	"table-export/config"
	"table-export/define"
	"table-export/export"
	"table-export/meta"
)

func main() {

	log.SetReportCaller(true)
	log.SetFormatter(&Formatter{})
	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	log.SetLevel(log.InfoLevel)

	flag.Parse()

	if params.help {
		usage()
		return
	}

	if params.version {
		define.PrintBuild()
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
