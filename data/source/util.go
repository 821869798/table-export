package source

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/meta"
)

func createOrAddDataModel(tableSource *meta.TableSource, dataModel *model.TableModel, rowData [][]string) error {
	//数据行数不够
	if len(rowData) < config.GlobalConfig.Table.DataStart {
		return errors.New(fmt.Sprintf("excel source file[%v] sheet[%v] row count must >= %v", tableSource.Table, tableSource.Sheet, config.GlobalConfig.Table.DataStart))
	}

	//名字行，先存成一个map
	nameRow := rowData[config.GlobalConfig.Table.Name]
	nameIndex := make(map[string]int)
	for index, name := range nameRow {
		//先判断是否需要这个字段
		if _, ok := dataModel.Meta.SourceMap[name]; !ok {
			continue
		}
		//excel中需要的字段重复了
		if _, ok := nameIndex[name]; ok {
			return errors.New(fmt.Sprintf("excel source file[%v] sheet[%v] name repeated[%v]", tableSource.Table, tableSource.Sheet))
		}
		nameIndex[name] = index
	}

	if dataModel.NameIndexMapping == nil {
		dataModel.NameIndexMapping = make(map[string]int)
		//只有一个的情况
		for _, tf := range dataModel.Meta.Fields {
			sourceIndex, ok := nameIndex[tf.Source]
			if !ok {
				return errors.New(fmt.Sprintf("excel source file[%v] sheet[%v] can't find field name[%v]", tableSource.Table, tableSource.Sheet, tf.Source))
			}
			dataModel.NameIndexMapping[tf.Target] = sourceIndex
		}
		dataModel.NameRow = nameRow
		dataModel.DescRow = rowData[config.GlobalConfig.Table.Desc]
		dataModel.RawData = rowData[config.GlobalConfig.Table.DataStart:]
	} else {
		//检验多分配置的索引以及顺序是否一致
		for _, tf := range dataModel.Meta.Fields {
			sourceIndex, ok := nameIndex[tf.Source]
			if !ok {
				return errors.New(fmt.Sprintf("excel source file[%v] sheet[%v] can't find field name[%v]", tableSource.Table, tableSource.Sheet, tf.Source))
			}
			oldIndex, ok := dataModel.NameIndexMapping[tf.Target]
			if sourceIndex != oldIndex {
				return errors.New(fmt.Sprintf("excel multiply sources field name not inconformity,target:%v", dataModel.Meta.Target))
			}
		}
		dataModel.RawData = append(dataModel.RawData, rowData[config.GlobalConfig.Table.DataStart:]...)
	}

	return nil
}
