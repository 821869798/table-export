package export

import (
	_ "github.com/360EntSecGroup-Skylar/excelize/v2"
	log "github.com/sirupsen/logrus"
	"strings"
	"table-export/config"
	"table-export/meta"
)

type Entry struct {
	mode string
}

func NewEntry(mode string) *Entry {
	e := &Entry{
		mode: mode,
	}
	return e
}

func (e *Entry) Run() {
	modeSlice := strings.Split(e.mode, ",")
	for _, mode := range modeSlice {
		metaRule, ok := config.GlobalConfig.Meta.Rule[mode]
		if !ok {
			log.WithFields(log.Fields{
				"mode": mode,
			}).Error("find meta rule failed")
			continue
		}

		tableMetas, err := meta.LoadTableMetasByDir(metaRule.ConfigDir)
		if err != nil {
			log.WithFields(log.Fields{
				"mode": mode,
				"err":  err,
			}).Fatal("load table meta toml config failed")
		}

		if creatorFunc, ok := exportCreators[metaRule.ExportType]; ok {

			log.WithFields(log.Fields{
				"mode": mode,
			}).Debug("start run export")

			export := creatorFunc(tableMetas)
			go func() {
				export.Export()
			}()
		} else {
			log.Fatalf("export mode can't support:%v", metaRule.ExportType)
		}
	}
}
