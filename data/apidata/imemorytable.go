package apidata

type IMemoryTable interface {
	TableName() string
	RawDataMapping() map[interface{}]interface{}
	RawDataList() []map[string]interface{}
	GetRecordByRow(rowIndex int) (map[string]interface{}, error)
	GetRecordRecordMap(recordIndex int) map[string]interface{}
	GetRecordByKey(keys ...interface{}) map[string]interface{}
	RowIndexList() []int
	AddExtraData(string, interface{})
	RemoveExtraData(string)
	GetExtraDataMap() map[string]interface{}
}
