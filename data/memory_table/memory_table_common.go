package memory_table

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/wrap"
	"github.com/821869798/table-export/data/apidata"
	"github.com/821869798/table-export/data/model"
)

// MemoryTableCommon Table中所有的key的类型有不一样的,只能使用interface{}
type MemoryTableCommon struct {
	ValueMapping map[interface{}]interface{}
	ValueList    []map[string]interface{}
	RowDataIndex []int // 数据所在的原始行，可能之后会支持删除某一行
	Name         string
}

func NewMemTableCommon(dataModel *model.TableModel, count, cap int) (apidata.IMemoryTable, error) {
	// count是有效数据,cap是容量，兼容以后可能允许空行的情况
	memTable := &MemoryTableCommon{
		ValueMapping: make(map[interface{}]interface{}, count),
		ValueList:    make([]map[string]interface{}, 0, count),
		RowDataIndex: make([]int, cap),
		Name:         dataModel.Meta.Target,
	}
	err := memTable.ReadTableModel(dataModel)
	if err != nil {
		return nil, err
	}
	return memTable, nil
}

func (m *MemoryTableCommon) TableName() string {
	return m.Name
}

func (m *MemoryTableCommon) RawDataMapping() map[interface{}]interface{} {
	return m.ValueMapping
}

func (m *MemoryTableCommon) RawDataList() []map[string]interface{} {
	return m.ValueList
}

func (m *MemoryTableCommon) ReadTableModel(dataModel *model.TableModel) error {
	rowDataOffset := config.GlobalConfig.Table.DataStart + 1
	for rowIndex, rowData := range dataModel.RawData {
		//数据表中一行数据的字符串
		recordMap := make(map[string]interface{}, len(dataModel.Meta.Fields))
		keys := make([]interface{}, len(dataModel.Meta.Keys))
		for _, tf := range dataModel.Meta.Fields {
			rawIndex := dataModel.NameIndexMapping[tf.Target]
			var rawStr string
			if rawIndex < len(rowData) {
				rawStr = rowData[rawIndex]
			}
			output, err := wrap.GetOutputValue(config.ExportType_Json, tf.Type, rawStr)
			if err != nil {
				return errors.New(fmt.Sprintf("create memory table failed,file[%v] RowCount[%v] filedName[%v] error:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source, err))
			}

			recordMap[tf.Target] = output

			//存储key
			if tf.Key > 0 {
				keys[tf.Key-1] = output
			}
		}

		//写入key和数据
		dataMap := m.ValueMapping
		for i := 0; i < len(keys)-1; i++ {
			keyStr := keys[i]
			tmpInterface, ok := dataMap[keyStr]
			var tmpMap map[interface{}]interface{}
			if !ok {
				tmpMap = make(map[interface{}]interface{})
				dataMap[keyStr] = tmpMap
			} else {
				tmpMap = tmpInterface.(map[interface{}]interface{})
			}
			dataMap = tmpMap
		}

		//判断唯一key是否重复
		lastKey := keys[len(keys)-1]
		_, ok := dataMap[lastKey]
		if ok {
			return errors.New(fmt.Sprintf("create memory table failed,file[%v] RowCount[%v] key is repeated:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, keys))
		}

		dataMap[lastKey] = recordMap
		m.ValueList = append(m.ValueList, recordMap)

		// 这里没有-1是为了避免为0,把0当作无效数据，不然没法区分
		m.RowDataIndex[rowIndex] = len(m.ValueList)
	}

	return nil
}

func (m *MemoryTableCommon) GetRecordByRow(rowIndex int) (map[string]interface{}, error) {
	recordIndex := m.RowDataIndex[rowIndex]
	if recordIndex > 0 {
		return m.ValueList[recordIndex-1], nil
	}
	return nil, errors.New(fmt.Sprintf("rowIndex[%v] is not exist", rowIndex))
}

func (m *MemoryTableCommon) GetRecordRecordMap(recordIndex int) map[string]interface{} {
	return m.ValueList[recordIndex-1]
}

func (m *MemoryTableCommon) RowIndexList() []int {
	return m.RowDataIndex
}
