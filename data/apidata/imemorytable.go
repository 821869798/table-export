package apidata

type IMemoryTable interface {
	TableName() string
	GetRecordByRow(rowIndex int) (map[string]interface{}, error)
	GetRecordRecordMap(recordIndex int) map[string]interface{}
	RowIndexList() []int
}
