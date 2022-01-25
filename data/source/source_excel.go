package source

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"path/filepath"
	"table-export/config"
	"table-export/data/api"
	"table-export/data/model"
	"table-export/meta"
)

type DataSourceExcel struct {
}

func NewDataSourceExcel() api.DataSource {
	d := &DataSourceExcel{}
	return d
}

func (d *DataSourceExcel) LoadDataModel(tableMetal *meta.TableMeta) (*model.TableModel, error) {
	//检测数量
	if len(tableMetal.Sources) == 0 {
		return nil, errors.New(fmt.Sprintf("table meta config[%v] sources count must >= 0", tableMetal.Target))
	}

	//读取excel数据到自定义结构体中
	dataModel := model.NewTableModel(tableMetal)
	for _, tableSource := range tableMetal.Sources {
		filePath := filepath.Join(config.GlobalConfig.Table.SrcDir, tableSource.Table)
		excelFile, err := excelize.OpenFile(filePath)
		if err != nil {
			return nil, err
		}
		rowData, err := excelFile.GetRows(tableSource.Sheet)
		if err != nil {
			return nil, err
		}
		// 创建或者添加数据
		err = createOrAddDataModel(tableSource, dataModel, rowData)
		if err != nil {
			return nil, err
		}
	}
	return dataModel, nil
}