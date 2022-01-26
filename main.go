package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"table-export/config"
	"table-export/define"
	"table-export/export"
	"table-export/meta"
)

func main() {

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		//以下设置只是为了使输出更美观
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:03:04",
	})
	log.SetOutput(os.Stdout)
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
