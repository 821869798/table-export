package source

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"table-export/config"
	"table-export/data/api"
	"table-export/data/model"
	"table-export/meta"
)

type DataSourceCsv struct {
}

func NewDataSourceCsv() api.DataSource {
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
		filePath := config.AbsExeDir(config.GlobalConfig.Table.SrcDir, tableSource.Table)

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

		// 创建或者添加数据
		err = createOrAddDataModel(tableSource, dataModel, rowData)
		if err != nil {
			return nil, err
		}
	}
	return dataModel, nil
}
