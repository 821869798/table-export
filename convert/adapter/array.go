package adapter

import "table-export/meta"

type Array struct {
	Datas     []string
	ValueType *meta.TableFieldType
}

func NewArray(datas []string, fieldType *meta.TableFieldType) *Array {
	a := &Array{
		Datas:     datas,
		ValueType: fieldType,
	}
	return a
}

var EmptyArray *Array = NewArray(nil, meta.TableFieldTypeNone)
