package json

import (
	"encoding/json"
	"github.com/gookit/slog"
	"os"
	"path/filepath"
	"table-export/config"
	"table-export/convert/wrap"
	"table-export/data/model"
	"table-export/export/api"
	"table-export/export/common"
	"table-export/meta"
	"table-export/util"
	"time"
)

type ExportJson struct {
	tableMetas []*meta.RawTableMeta
}

func NewExportJson(tableMetas []*meta.RawTableMeta, extraArg map[string]string) api.IExport {
	e := &ExportJson{
		tableMetas: tableMetas,
	}
	return e
}

func (e *ExportJson) TableMetas() []*meta.RawTableMeta {
	return e.tableMetas
}

func (e *ExportJson) Export(ru config.MetaRuleUnit) {

	jsonRule, ok := ru.(*config.RawMetaRuleUnitJson)
	if !ok {
		slog.Fatal("Export Json expect *RawMetaRuleUnitJson Rule Unit")
	}

	outputPath := config.AbsExeDir(jsonRule.JsonOutputDir)
	//清空目录
	if err := util.InitDirAndClearFile(outputPath, `^.*?\.json$`); err != nil {
		slog.Fatal(err)
	}

	defer util.TimeCost(time.Now(), "export json time cost = %v\n")

	//实际开始转换
	common.CommonMutilExport(e.tableMetas, func(dataModel *model.TableModel) {
		ExportJsonFile(dataModel, outputPath)
	})

}

func ExportJsonFile(dataModel *model.TableModel, outputPath string) {

	writeContent, _ := generateJsonData(dataModel, false)

	exportJson(dataModel, outputPath, writeContent)
}

func ExportListJsonFile(dataModel *model.TableModel, outputPath string) {
	writeContent, writeList := generateJsonData(dataModel, true)

	writeRoot := map[string]interface{}{
		"dataMap":  writeContent,
		"dataList": writeList,
	}

	exportJson(dataModel, outputPath, writeRoot)
}

func generateJsonData(dataModel *model.TableModel, exportList bool) (map[string]interface{}, []interface{}) {
	//创建存储的数据结构
	writeContent := make(map[string]interface{})
	var dataList []interface{} = nil
	// 导出的是列表模式
	if exportList {
		dataList = make([]interface{}, 0)
	}
	rowDataOffset := config.GlobalConfig.Table.DataStart + 1
	for rowIndex, rowData := range dataModel.RawData {
		//数据表中一行数据的字符串
		recordMap := make(map[string]interface{}, len(dataModel.Meta.Fields))
		keys := make([]string, len(dataModel.Meta.Keys))
		for _, tf := range dataModel.Meta.Fields {
			rawIndex := dataModel.NameIndexMapping[tf.Target]
			var rawStr string
			if rawIndex < len(rowData) {
				rawStr = rowData[rawIndex]
			}
			output, err := wrap.GetOutputValue(config.ExportType_Json, tf.Type, rawStr)
			if err != nil {
				slog.Fatalf("export json target file[%v] RowCount[%v] filedName[%v] error:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source, err)
			}

			recordMap[tf.Target] = output

			//存储key
			if tf.Key > 0 {
				formatKey, err := wrap.GetOutputStringValue(config.ExportType_Json, tf.Type, output)
				if err != nil {
					slog.Fatalf("export json target file[%v] RowCount[%v] filedName[%v] format key error:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source, err)
				}
				keys[tf.Key-1] = formatKey
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
			slog.Fatalf("export json target file[%v] RowCount[%v] key is repeated:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, keys)
		}

		if exportList {
			dataMap[lastKey] = len(dataList)
			dataList = append(dataList, recordMap)
		} else {
			dataMap[lastKey] = recordMap
		}
	}
	return writeContent, dataList
}

func exportJson(dataModel *model.TableModel, outputPath string, writeContent any) {
	b, err := json.MarshalIndent(writeContent, "", "    ")
	if err != nil {
		slog.Fatalf("export json error:%v", err)
	}

	filePath := filepath.Join(outputPath, dataModel.Meta.Target+".json")

	f, err := os.Create(filePath) //os.OpenFile(destFile, os.O_WRONLY, 0600)
	if err != nil {
		slog.Fatalf("export json error:%v", err)
	}

	_, err = f.Write(b)
	if err != nil {
		slog.Fatalf("export json error:%v", err)
	}

	err = f.Close()
	if err != nil {
		slog.Fatalf("export json error:%v", err)
	}
}
