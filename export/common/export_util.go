package common

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"table-export/config"
	"table-export/data"
	"table-export/data/model"
	"table-export/define"
	"table-export/export/wrap"
	"table-export/meta"
)

//通用的并行执行的方法a
func CommonMutilExport(tableMetas []*meta.RawTableMeta, exportFunc func(*model.TableModel)) {
	wg := sync.WaitGroup{}
	wg.Add(len(tableMetas))
	for _, tableMeta := range tableMetas {
		if tableMeta.Mode == define.CommentSymbol {
			//是注释模式，不导出
			wg.Done()
			continue
		}
		go func(tableMeta *meta.RawTableMeta) {
			tm, err := meta.NewTableMeta(tableMeta)
			if err != nil {
				log.Fatal(err)
			}

			dataModel, err := data.GetDataModelByType(tm)
			if err != nil {
				log.Fatal(err)
			}

			//执行函数
			exportFunc(dataModel)

			wg.Done()
		}(tableMeta)
	}
	wg.Wait()
}

func CommonGetWriteContentMap(dataModel *model.TableModel) map[string]interface{} {
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
				log.Fatalf("export target file[%v] RowCount[%v] filedName[%v] error:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source, err.Error())
			}
			outputStr, ok := output.(string)
			if !ok {
				log.Fatalf("export target file[%v] RowCount[%v] filedName[%v] convert to string error", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source)
			}
			recordString += tf.Target + "=" + outputStr + ","

			//存储key
			if tf.Key > 0 {
				if outputStr == "" {
					log.Fatalf("export target file[%v] RowCount[%v] filedName[%v] key content is null", dataModel.Meta.Target, rowIndex+rowDataOffset)
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
			log.Fatalf("export target file[%v] RowCount[%v] key is repeated:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, keys)
		}
		dataMap[lastKey] = recordString

	}
	return writeContent
}
