package lua

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"table-export/config"
	"table-export/data/model"
	"table-export/define"
	"table-export/export/api"
	"table-export/export/common"
	"table-export/export/wrap"
	"table-export/meta"
	"table-export/util"
	"text/template"
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

func (e *ExportLua) Export() {

	luaRule := config.GlobalConfig.Meta.RuleLua

	//清空目录
	if err := util.ClearDirAndCreateNew(luaRule.TempDir); err != nil {
		log.Fatal(err)
	}

	//实际开始转换
	common.CommonMutilExport(e.tableMetas, func(dataModel *model.TableModel) {
		exportLuaFile(dataModel, luaRule.TempDir)
	})

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
			output, err := wrap.GetOutputValue(define.ExportType_Lua, tf.Type, rawStr)
			if err != nil {
				log.Fatalf("export lua target file[%v] RowCount[%v] filedName[%v] error:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source, err.Error())
			}
			outputStr, ok := output.(string)
			if !ok {
				log.Fatalf("export lua target file[%v] RowCount[%v] filedName[%v] convert to string error", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source)
			}
			recordString += tf.Target + "=" + outputStr + ","

			//存储key
			if tf.Key > 0 {
				if outputStr == "" {
					log.Fatalf("export lua target file[%v] RowCount[%v] filedName[%v] key content is null", dataModel.Meta.Target, rowIndex+rowDataOffset)
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
			log.Fatalf("export lua target file[%v] RowCount[%v] key is repeated:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, keys)
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
		log.Fatal(err)
	}

	//填充数据
	result := &luaWriteContent{
		Name:    dataModel.Meta.Target,
		Content: writeContent,
	}

	//渲染输出
	err = tmpl.Execute(file, result)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}
