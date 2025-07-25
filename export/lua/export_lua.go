package lua

import (
	"github.com/821869798/fankit/fanpath"
	"github.com/821869798/fankit/fansys"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/wrap"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/export/api"
	"github.com/821869798/table-export/export/common"
	"github.com/821869798/table-export/meta"
	"github.com/821869798/table-export/util"
	"github.com/gookit/slog"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"
	"time"
)

type ExportLua struct {
	tableMetas []*meta.RawTableMeta
}

func NewExportLua(tableMetas []*meta.RawTableMeta, extraArg map[string]string) api.IExport {
	e := &ExportLua{
		tableMetas: tableMetas,
	}
	return e
}

func (e *ExportLua) TableMetas() []*meta.RawTableMeta {
	return e.tableMetas
}

func (e *ExportLua) Export(ru config.MetaRuleUnit) {

	luaRule, ok := ru.(*config.RawMetaRuleUnitLua)
	if !ok {
		slog.Fatal("Export Lua expect *RawMetaRuleUnitLua Rule Unit")
	}

	//清空目录
	if luaRule.EnableProcess {
		if err := fanpath.ClearDirAndCreateNew(fanpath.RelExecuteDir(luaRule.TempDir)); err != nil {
			slog.Fatal(err)
		}
	}

	if err := fanpath.InitDirAndClearFile(fanpath.RelExecuteDir(luaRule.LuaOutputDir), `^.*?\.(lua|meta)$`); err != nil {
		slog.Fatal(err)
	}

	defer util.TimeCost(time.Now(), "export lua time cost = %v\n")

	var outputPath string
	if luaRule.EnableProcess {
		outputPath = fanpath.RelExecuteDir(luaRule.TempDir)
	} else {
		outputPath = fanpath.RelExecuteDir(luaRule.LuaOutputDir)
	}

	//实际开始转换
	common.LoadTableModelParallel(e.tableMetas, func(dataModel *model.TableModel) {
		exportLuaFile(dataModel, outputPath)
	})

	if luaRule.EnableProcess {
		luaTablePostProcess(luaRule)
		_ = os.RemoveAll(fanpath.RelExecuteDir(luaRule.TempDir))
	}

}

// lua打表后处理，例如做优化，提前做一些表结构变化等
func luaTablePostProcess(luaRule *config.RawMetaRuleUnitLua) {
	//做lua导出后的预处理
	var execPath string
	if runtime.GOOS == "windows" {
		execPath = fanpath.RelExecuteDir(luaRule.LuaWinDir)
	} else {
		execPath = fanpath.RelExecuteDir(luaRule.LuaMacDir)
	}

	// 添加执行权限
	if output, err := fansys.AddExecuteChmod(execPath); nil != err {
		slog.Fatalf("Add lua execute chmod linux or macos exec error:%s\nerr:%v", output, err)
		os.Exit(1)
	}

	luaTempDir := fanpath.RelExecuteDir(luaRule.TempDir)
	outputDir := fanpath.RelExecuteDir(luaRule.LuaOutputDir)

	postProcessExec := exec.Command(
		execPath,
		fanpath.RelExecuteDir(luaRule.PostProcessLua),
		fanpath.RelExecuteDir(luaRule.PostWorkDir),
		luaTempDir,
		runtime.GOOS,
	)

	if output, err := postProcessExec.CombinedOutput(); nil != err {
		slog.Fatalf("lua postprocess exec error:%s\nerr:%v", output, err)
		os.Exit(1)
	}
	fileList, err := fanpath.GetFileListByExt(luaTempDir, ".lua")
	if err != nil {
		slog.Fatal(err)
	}

	//拷贝到目标去
	for _, file := range fileList {
		src, err := os.Open(file)
		if nil != err {
			slog.Fatal(err)
		}

		dst, err := os.Create(filepath.Join(outputDir, filepath.Base(file)))
		if err != nil {
			_ = src.Close()
			slog.Fatal(err)
		}
		_, err = io.Copy(dst, src)

		_ = src.Close()
		_ = dst.Close()

		if err != nil {
			slog.Fatal(err)
		}
	}
}

type luaWriteContent struct {
	Name    string
	Content interface{}
}

func exportLuaFile(dataModel *model.TableModel, outputPath string) {
	//创建存储的数据结构
	writeContent := make(map[string]interface{})
	rowDataOffset := config.GlobalConfig.Table.DataStart + 1
	for rowIndex, rowData := range dataModel.RawData {
		//数据表中一行数据的字符串
		recordString := ""
		keys := make([]string, len(dataModel.Meta.Keys))
		for _, tf := range dataModel.Meta.Fields {
			rawIndex := dataModel.NameIndexMapping[tf.Target]
			var rawStr string
			if rawIndex < len(rowData) {
				rawStr = rowData[rawIndex]
			}
			outputStr, err := wrap.GetOutputStringValue(config.ExportType_Lua, tf.Type, rawStr)
			if err != nil {
				slog.Fatalf("export lua target file[%v] RowCount[%v] filedName[%v] error:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source, err.Error())
			}
			recordString += tf.Target + "=" + outputStr + ","

			//存储key
			if tf.Key > 0 {
				if outputStr == "" {
					slog.Fatalf("export lua target file[%v] RowCount[%v] filedName[%v] key content is null", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source)
				}
				keys[tf.Key-1] = outputStr
			}
		}

		//写入key和数据
		dataMap := writeContent
		for i := 0; i < len(keys)-1; i++ {
			keyStr := keys[i]
			tmpInterface, ok := dataMap[keyStr]
			var tmpMap map[string]interface{}
			if !ok {
				tmpMap = make(map[string]interface{})
				dataMap[keyStr] = tmpMap
			} else {
				tmpMap = tmpInterface.(map[string]interface{})
			}
			dataMap = tmpMap
		}

		//判断唯一key是否重复
		lastKey := keys[len(keys)-1]
		_, ok := dataMap[lastKey]
		if ok {
			slog.Fatalf("export lua target file[%v] RowCount[%v] key is repeated:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, keys)
		}
		dataMap[lastKey] = recordString
	}

	//创建模板,绑定全局函数,并且解析
	tmpl, err := template.New("export_lua").Funcs(template.FuncMap{
		"IsString": func(v interface{}) bool {
			_, ok := v.(string)
			return ok
		},
		"ToString": func(v interface{}) string {
			str, ok := v.(string)
			if !ok {
				return ""
			}
			return str
		},
	}).Parse(templateLua)

	filePath := filepath.Join(outputPath, dataModel.Meta.Target+".lua")
	file, err := os.Create(filePath)
	if err != nil {
		slog.Fatal(err)
	}

	//填充数据
	result := &luaWriteContent{
		Name:    dataModel.Meta.Target,
		Content: writeContent,
	}

	//渲染输出
	err = tmpl.Execute(file, result)
	if err != nil {
		slog.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		slog.Fatal(err)
	}

}
