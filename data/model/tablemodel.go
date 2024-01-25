package model

import (
	"github.com/821869798/table-export/data/apidata"
	"github.com/821869798/table-export/meta"
)

// TableModel 表的数据
type TableModel struct {
	NameIndexMapping map[string]int //目标字段名到index的索引
	NameRow          []string       //字段名的行
	DescRow          []string       //描述的行
	RawData          [][]string     //数据
	Meta             *meta.TableMeta
	Optimize         *TableOptimize
	MemTable         apidata.IMemoryTable // 内存中转换后的数据，主要是interface{}类型
	ClearRecord      bool                 // 是否清除整个表的每行数据
}

func NewTableModel(meta *meta.TableMeta) *TableModel {
	d := &TableModel{
		NameIndexMapping: nil,
		NameRow:          nil,
		DescRow:          nil,
		RawData:          nil,
		Meta:             meta,
		ClearRecord:      false,
	}
	return d
}
