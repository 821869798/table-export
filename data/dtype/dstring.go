package dtype

import (
	"fmt"
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
)

var (
	dstringDefault = &DStringType{
		value: "",
	}
)

type DStringType struct {
	value string
}

func (ds *DStringType) String() string {
	return fmt.Sprintf("<DString:%s>", ds.value)
}

func (ds *DStringType) Value() string {
	return ds.value
}

func (ds *DStringType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {
	//TODO implement me
	panic("implement me")
}

func (ds *DStringType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {
	//TODO implement me
	panic("implement me")
}

func (ds *DStringType) IsEmptyOrZero() bool {
	return ds.value == ""
}

func NewDString(value string) *DStringType {
	if value == "" {
		return dstringDefault
	}
	return &DStringType{value: value}
}
