package json

import (
	"encoding/json"
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

func (e *ExportJson) Export() {
	jsonRule := config.GlobalConfig.Meta.RuleJson

	//清空目录
	if err := util.ClearDirAndCreateNew(jsonRule.OutputDir); err != nil {
		log.Fatal(err)
	}

	//实际开始转换
	common.CommonMutilExport(e.tableMetas, func(dataModel *model.TableModel) {
		exportJsonFile(dataModel, jsonRule.OutputDir)
	})

}

func exportJsonFile(dataModel *model.TableModel, outputPath string) {
	//创建存储的数据结构
	writeContent := make(map[string]interface{})
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
			output, err := wrap.GetOutputValue(define.ExportType_Json, tf.Type, rawStr)
			if err != nil {
				log.Fatalf("export json target file[%v] RowCount[%v] filedName[%v] error:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source, err.Error())
			}

			recordMap[tf.Target] = output

			//存储key
			if tf.Key > 0 {
				formatKey, err := wrap.GetFormatValue(define.ExportType_Json, tf.Type, output)
				if err != nil {
					log.Fatalf("export json target file[%v] RowCount[%v] filedName[%v] format key error:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source, err.Error())
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
			log.Fatalf("export json target file[%v] RowCount[%v] key is repeated:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, keys)
		}
		dataMap[lastKey] = recordMap
	}

	b, err := json.MarshalIndent(writeContent, "", "    ")
	if err != nil {
		log.Fatalf("export json error:%v", err)
	}

	filePath := filepath.Join(outputPath, dataModel.Meta.Target+".json")

	f, err := os.Create(filePath) //os.OpenFile(destFile, os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("export json error:%v", err)
	}

	_, err = f.Write(b)
	if err != nil {
		log.Fatalf("export json error:%v", err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalf("export json error:%v", err)
	}
}
