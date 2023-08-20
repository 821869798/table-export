package dtype

import (
	"fmt"
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
	"math"
)

var (
	ddoubleDefault = &DDoubleType{
		value: 0,
	}
)

type DDoubleType struct {
	value float64
}

func (dd *DDoubleType) String() string {
	return fmt.Sprintf("<DDouble:%f>", dd.value)
}

func (dd *DDoubleType) Value() float64 {
	return dd.value
}

func (dd *DDoubleType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (dd *DDoubleType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (dd *DDoubleType) IsEmptyOrZero() bool {
	return math.Abs(dd.value) < epsilon
}

func NewDDouble(value float64) *DDoubleType {
	if math.Abs(value) < epsilon {
		return ddoubleDefault
	}
	return &DDoubleType{value: value}
}
