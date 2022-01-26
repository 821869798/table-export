package export

import (
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
	"table-export/config"
	"table-export/define"
	"table-export/meta"
)

type Entry struct {
	mode  string //转换模式
	extra string //额外参数
}

func NewEntry(mode, extra string) *Entry {
	e := &Entry{
		mode:  mode,
		extra: extra,
	}
	return e
}

func (e *Entry) Run() {
	extraArg := make(map[string]string)
	strMap1 := strings.Split(e.extra, "|")
	for _, v := range strMap1 {
		if v == "" {
			continue
		}
		kvStr := strings.Split(v, "=")
		if len(kvStr) == 2 {
			extraArg[kvStr[0]] = kvStr[1]
		}
	}

	wg := sync.WaitGroup{}

	modeSlice := strings.Split(e.mode, "|")
	for _, mode := range modeSlice {
		exportType, ok := define.GetExportTypeFromString(mode)
		if !ok {
			log.Fatalf("export mode can't support:%v", mode)
		}
		metaRule := config.GlobalConfig.Meta.GetRawMetaBaseConfig(exportType)

		tableMetas, err := meta.LoadTableMetasByDir(config.AbsExeDir(metaRule.ConfigDir))
		if err != nil {
			log.WithFields(log.Fields{
				"mode": mode,
				"err":  err,
			}).Fatal("load table meta toml config failed")
		}

		if creatorFunc, ok := exportCreators[exportType]; ok {

			log.WithFields(log.Fields{
				"mode": mode,
			}).Debug("start run export")

			export := creatorFunc(tableMetas, extraArg)
			wg.Add(1)
			go func() {
				export.Export()
				wg.Done()
			}()

		} else {
			log.Fatalf("export mode can't support:%v", exportType)
		}
	}

	wg.Wait()
}
