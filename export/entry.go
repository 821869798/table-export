package export

import (
	"github.com/821869798/fankit/fanpath"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/data/env"
	"github.com/821869798/table-export/meta"
	"github.com/gookit/slog"
	"strings"
	"sync"
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
		} else if len(kvStr) == 1 {
			extraArg[kvStr[0]] = ""
		}
	}

	wg := sync.WaitGroup{}

	modeSlice := strings.Split(e.mode, "|")
	for _, mode := range modeSlice {
		metaRule := config.GetMetaRuleConfigByName(mode)
		if metaRule == nil {
			slog.Fatalf("export mode can't not find in config:%v", mode)
		}

		tableMetas, err := meta.LoadTableMetasByDir(fanpath.RelExecuteDir(metaRule.ConfigDir))
		if err != nil {
			slog.Fatalf("load table meta toml config failed! mode:%s err:%v", mode, err)
		}

		for _, rule := range metaRule.RuleUnits {
			if creatorFunc, ok := exportCreators[rule.RuleExportType()]; ok {
				slog.Debugf("start run export mode:%s", mode)

				export := creatorFunc(tableMetas, extraArg)
				//因为有个全局的环境，所以不支持同时转换多个mode，需要依次执行
				env.InitEnv()
				if rulePlus, ok := rule.(config.MetaRuleUnitPlus); ok {
					env.SetMetaRuleUnitPlus(rulePlus)
				}
				export.Export(rule)
				//wg.Add(1)
				//go func() {
				//	export.Export(rule)
				//	wg.Done()
				//}()
			} else {
				slog.Fatalf("export mode can't support:%v", rule.RuleExportType())
			}
		}

	}

	wg.Wait()
}
