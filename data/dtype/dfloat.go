package dtype

import (
	"fmt"
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
	"math"
)

var (
	dfloatDefault = &DFloatType{
		value: 0,
	}
)

type DFloatType struct {
	value float32
}

func (df *DFloatType) String() string {
	return fmt.Sprintf("<DFloat:%f>", df.value)
}

func (df *DFloatType) Value() float32 {
	return df.value
}

func (df *DFloatType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (df *DFloatType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (df *DFloatType) IsEmptyOrZero() bool {
	const epsilon = 1e-9
	return math.Abs(float64(df.value)) < epsilon
}

func NewDFloat(value float32) *DFloatType {
	if math.Abs(float64(value)) < epsilon {
		return dfloatDefault
	}
	return &DFloatType{value: value}
}
