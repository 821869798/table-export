package model

import "github.com/821869798/table-export/meta"

type TableOptimize struct {
	OptimizeFields []*TableOptimizeField
	FieldsMap      map[string]int
}

func NewTableOptimize() *TableOptimize {
	t := &TableOptimize{
		FieldsMap: make(map[string]int, 0),
	}
	return t
}

func (t *TableOptimize) AddOptimizeField(field *TableOptimizeField) {
	t.FieldsMap[field.Field.Target] = len(t.OptimizeFields)
	t.OptimizeFields = append(t.OptimizeFields, field)
}

func (t *TableOptimize) GetOptimizeField(field *meta.TableField) (*TableOptimizeField, int) {
	fIndex, ok := t.FieldsMap[field.Target]
	if !ok {
		return nil, -1
	}

	return t.OptimizeFields[fIndex], fIndex
}

type TableOptimizeField struct {
	Field        *meta.TableField
	OriginDatas  []string
	DataIndexs   []int
	OptimizeType *meta.TableFieldType
}

func NewTableOptimizeField(field *meta.TableField, valueCount int, allCount int) *TableOptimizeField {
	t := &TableOptimizeField{
		Field:        field,
		OriginDatas:  make([]string, valueCount),
		DataIndexs:   make([]int, allCount),
		OptimizeType: field.Type.CreateArrayFieldType(),
	}
	return t
}
