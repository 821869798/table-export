package data

import (
	"github.com/821869798/table-export/data/source"
)

var dataSourceLoads map[string]source.IDataSource

func init() {
	dataSourceLoads = make(map[string]source.IDataSource)
	dataSourceLoads["excel"] = source.NewDataSourceExcel()
	dataSourceLoads["csv"] = source.NewDataSourceCsv()
}
