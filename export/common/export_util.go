package common

import (
	"github.com/821869798/fankit/fanpath"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/data/check"
	"github.com/821869798/table-export/data/env"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/ext"
	"github.com/821869798/table-export/ext/ext_field_script"
	"github.com/821869798/table-export/ext/ext_post"
	"github.com/821869798/table-export/meta"
	"github.com/BurntSushi/toml"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/gookit/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func ExportPlusCommon(tableMetas []*meta.RawTableMeta, rulePlus config.MetaRuleUnitPlus) []*model.TableModel {

	// 加载枚举配置
	var enumFiles = rulePlus.GetEnumFiles()
	rawEnumConfigs := make([]*config.RawMetaEnumConfig, 0, len(enumFiles))
	for _, p := range enumFiles {
		matches, err := doublestar.FilepathGlob(p)
		if err != nil {
			slog.Fatalf("Enum Files load error filePath:%s err:%v", p, err)
		}
		for _, m := range matches {
			fullPath := fanpath.AbsOrRelExecutePath(m)
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

	// 添加内置扩展类型
	for _, extFieldTypeName := range rulePlus.GetBuiltinFieldTypes() {
		extFieldType, ok := ext.GetExistExtFieldType(extFieldTypeName)
		if !ok {
			slog.Fatalf("no builtin ext field type:%v", extFieldTypeName)
			os.Exit(1)
		}
		err := env.AddExtFieldType(extFieldType)
		if err != nil {
			slog.Fatalf("add ext field type error:%v", err)
			os.Exit(1)
		}
	}

	// 添加脚本实现的扩展类型
	for _, extFieldScriptPath := range rulePlus.GetExtFieldTypeScriptPath() {
		matches, err := doublestar.FilepathGlob(extFieldScriptPath)
		if err != nil {
			slog.Fatalf("Ext Script path glob error filePath:%s err:%v", extFieldScriptPath, err)
			os.Exit(1)
		}
		for _, m := range matches {
			fileExt := filepath.Ext(m)
			switch fileExt {
			case ".js":
				extFieldType, err := ext_field_script.NewExtFieldJS(m)
				if err != nil {
					slog.Fatalf("Ext Script Files laod error filePath:%s err:%v", m, err)
					os.Exit(1)
				}
				err = env.AddExtFieldType(extFieldType)
				if err != nil {
					slog.Fatalf("add ext field type error:%v", err)
					os.Exit(1)
				}
			default:
				slog.Fatalf("ext field script type not support:%s", fileExt)
			}
		}
	}

	//实际开始转换
	allDataModel := LoadTableModelPlusParallel(tableMetas, rulePlus, func(tableModel *model.TableModel) {
		// 单表的后处理
		scriptPath := strings.TrimSpace(tableModel.Meta.PostScript)
		if scriptPath != "" {
			fileExt := filepath.Ext(scriptPath)
			switch fileExt {
			case ".js":
				extPostTable, err := ext_post.NewExtPostTableJS(tableModel.Meta.PostScript)
				if err != nil {
					slog.Fatalf("Ext Script Files laod error filePath:%s err:%v", tableModel.Meta.PostScript, err)
					os.Exit(1)
				}
				err = extPostTable.PostTable(tableModel)
				if err != nil {
					slog.Fatalf("Ext Script post process error filePath:%s err:%v", tableModel.Meta.PostScript, err)
					os.Exit(1)
				}
			default:
				slog.Fatalf("ext post script type not support:%s", fileExt)
			}

		}

	})

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

	tableMap := make(map[string]*model.TableModel, len(allDataModel))
	for _, m := range allDataModel {
		tableMap[m.Meta.Target] = m
	}

	// 全局单次的后处理
	postScriptPath := strings.TrimSpace(rulePlus.GetPostGlobalScriptPath())
	if postScriptPath != "" {
		fileExt := filepath.Ext(postScriptPath)
		switch fileExt {
		case ".js":
			extPostGlobal, err := ext_post.NewExtPostGlobalJS(postScriptPath)
			if err != nil {
				slog.Fatalf("ExtPostGlobal Script Files laod error filePath:%s err:%v", postScriptPath, err)
				os.Exit(1)
			}
			err = extPostGlobal.PostGlobal(tableMap)
			if err != nil {
				slog.Fatalf("ExtPostGlobal Script post process error filePath:%s err:%v", postScriptPath, err)
				os.Exit(1)
			}
		default:
			slog.Fatalf("ext post global script type not support:%s", fileExt)
		}
	}

	return allDataModel
}
