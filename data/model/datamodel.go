package model

type DataModel struct {
	NameIndexMapping map[string]int //字段名到index的索引
	NameRow          []string       //字段名的行
	DescRow          []string       //描述的行
	RawData          [][]string     //数据
}

func NewDataModel() *DataModel {
	d := &DataModel{
		NameIndexMapping: make(map[string]int),
		NameRow:          nil,
		DescRow:          nil,
		RawData:          nil,
	}
	return d
}
