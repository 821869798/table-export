package excel

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

func (d *DataSourceExcel) LoadDataModel(tableMetal *meta.RawTableMeta) (*model.DataModel, error) {
	//for k := range tableMetal.Sources
	dataModel := model.NewDataModel()
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

func createOrAddDataModel(tableSource *meta.RawTableSource, dataModel *model.DataModel, rowData [][]string) error {
	//数据行数不够
	if len(rowData) < config.GlobalConfig.Table.DataStart {
		return errors.New(fmt.Sprintf("excel source file[%v] sheet[%v] row count must >= %v\n", tableSource.Table, tableSource.Sheet, config.GlobalConfig.Table.DataStart))
	}

	return nil
}
