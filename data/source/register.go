package source

var dataSourceLoads map[string]IDataSource

func init() {
	dataSourceLoads = make(map[string]IDataSource)
	dataSourceLoads["excel"] = NewDataSourceExcel()
	dataSourceLoads["csv"] = NewDataSourceCsv()
}
