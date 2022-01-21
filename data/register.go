package data

import (
	"table-export/data/api"
	"table-export/data/excel"
)

var dataSourceLoads map[string]api.DataSource

func init() {
	dataSourceLoads = make(map[string]api.DataSource)
	dataSourceLoads["excel"] = excel.NewDataSourceExcel()
}
