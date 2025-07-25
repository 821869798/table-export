package source

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/821869798/fankit/fanpath"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/meta"
	"os"
	"strings"
)

type DataSourceCsv struct {
}

func NewDataSourceCsv() IDataSource {
	d := &DataSourceCsv{}
	return d
}

func (d *DataSourceCsv) LoadDataModel(tableMetal *meta.TableMeta) (*model.TableModel, error) {
	//检测数量
	if len(tableMetal.Sources) == 0 {
		return nil, errors.New(fmt.Sprintf("table meta config[%v] sources count must >= 0", tableMetal.Target))
	}

	//读取excel数据到自定义结构体中
	dataModel := model.NewTableModel(tableMetal)
	for _, tableSource := range tableMetal.Sources {
		filePath := fanpath.RelExecuteDir(config.GlobalConfig.Table.SrcDir, tableSource.Table)

		csvFile, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		csvReader := csv.NewReader(csvFile) //创建一个新的写入文件流
		csvReader.LazyQuotes = true
		rowData, err := csvReader.ReadAll()

		if err != nil {
			return nil, err
		}

		err = csvFile.Close()
		if err != nil {
			return nil, err
		}

		//去除utf-8的bom
		if len(rowData) > 0 && len(rowData[0]) > 0 {
			firstData := rowData[0][0]
			rowData[0][0] = strings.TrimPrefix(firstData, "\uFEFF")
		}

		// 创建或者添加数据
		err = createOrAddDataModel(tableSource, dataModel, rowData)
		if err != nil {
			return nil, err
		}
	}
	return dataModel, nil
}
