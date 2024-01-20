package common

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/data/check"
	"github.com/821869798/table-export/data/env"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/meta"
	"github.com/821869798/table-export/util"
	"github.com/BurntSushi/toml"
	"github.com/gookit/slog"
	"path/filepath"
	"sync"
)

func ExportPlusCommon(tableMetas []*meta.RawTableMeta, rulePlus config.MetaRuleUnitPlus) []*model.TableModel {

	// 加载枚举配置
	var enumFiles = rulePlus.GetEnumFiles()
	rawEnumConfigs := make([]*config.RawMetaEnumConfig, 0, len(enumFiles))
	for _, p := range enumFiles {
		matches, err := filepath.Glob(p)
		if err != nil {
			slog.Fatalf("Enum Files laod error filePath:%s err:%v", p, err)
		}
		for _, m := range matches {
			fullPath := util.AbsOrRelExecutePath(m)
			enumConfig := new(config.RawMetaEnumConfig)
			if _, err := toml.DecodeFile(fullPath, enumConfig); err != nil {
				slog.Fatalf("load enum config error:%v", err)
			}
			rawEnumConfigs = append(rawEnumConfigs, enumConfig)
		}
	}
	if err := env.AddEnumDefines(rawEnumConfigs); err != nil {
		slog.Fatalf("add enum config error:%v", err)
	}

	// TODO 加载自定义解析脚本

	//实际开始转换
	allDataModel := LoadTableModelPlusParallel(tableMetas, rulePlus, nil)

	// TODO 表的数据后处理

	// 表的数据检查
	global := make(map[string]map[interface{}]interface{}, len(allDataModel))
	for _, m := range allDataModel {
		global[m.Meta.Target] = m.MemTable.RawDataMapping()
	}
	wgCheck := sync.WaitGroup{}
	wgCheck.Add(len(allDataModel))
	for _, m := range allDataModel {
		go func(m *model.TableModel) {
			check.Run(m, global)
			wgCheck.Done()
		}(m)
	}
	wgCheck.Wait()

	return allDataModel
}
