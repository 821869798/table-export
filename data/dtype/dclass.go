package dtype

import (
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
	"github.com/gookit/slog"
)

type DClassType struct {
	Fields     []apibase.DType
	FieldNames []string
	FieldsMap  map[string]apibase.DType
}

func (dc *DClassType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {

}

func (dc *DClassType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {

}

func (dc *DClassType) IsEmptyOrZero() bool {
	return false
}

func (dc *DClassType) AddField(name string, value apibase.DType) {
	_, ok := dc.FieldsMap[name]
	if ok {
		slog.Panicf("class add fields repeat! name :%s", name)
	}

	dc.Fields = append(dc.Fields, value)
	dc.FieldNames = append(dc.FieldNames, name)
	dc.FieldsMap[name] = value
}
