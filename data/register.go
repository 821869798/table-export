package data

import (
	"table-export/data/api"
	"table-export/data/source"
)

var dataSourceLoads map[string]api.DataSource

func init() {
	dataSourceLoads = make(map[string]api.DataSource)
	dataSourceLoads["excel"] = source.NewDataSourceExcel()
	dataSourceLoads["csv"] = source.NewDataSourceCsv()
}
