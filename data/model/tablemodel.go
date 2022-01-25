package model

import "table-export/meta"

//表的数据
type TableModel struct {
	NameIndexMapping map[string]int //目标字段名到index的索引
	NameRow          []string       //字段名的行
	DescRow          []string       //描述的行
	RawData          [][]string     //数据
	Meta             *meta.TableMeta
}

func NewTableModel(meta *meta.TableMeta) *TableModel {
	d := &TableModel{
		NameIndexMapping: nil,
		NameRow:          nil,
		DescRow:          nil,
		RawData:          nil,
		Meta:             meta,
	}
	return d
}
