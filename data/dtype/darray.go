package dtype

import (
	"fmt"
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
	"github.com/821869798/table-export/meta"
	"github.com/821869798/table-export/util"
)

type DArrayType struct {
	valueType *meta.TableFieldType
	dataArray []apibase.DType
}

func (da *DArrayType) String() string {
	return fmt.Sprintf("<DArray>[%s]", util.JoinEx(da.dataArray, ","))
}

func (da *DArrayType) Value() []apibase.DType {
	return da.dataArray
}

func (da *DArrayType) ValueType() *meta.TableFieldType {
	return da.valueType
}

func (da *DArrayType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (da *DArrayType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (da *DArrayType) IsEmptyOrZero() bool {
	return len(da.dataArray) == 0
}

func NewDArray(valueType *meta.TableFieldType, data []apibase.DType) *DArrayType {
	return &DArrayType{
		valueType: valueType,
		dataArray: data,
	}
}
